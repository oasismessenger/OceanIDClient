package IDClient

import (
	"context"
	"crypto/rand"
	"math/big"
	"sync/atomic"

	"github.com/pkg/errors"
	"google.golang.org/grpc"

	"github.com/oasismessenger/OceanIDClient/schemes/id_service"
)

type Counter struct {
	ctx    context.Context
	cancel context.CancelFunc

	limit int
	done  func() error

	C   chan struct{}
	Err chan error
}

func (c *Counter) Call() {
	c.C <- struct{}{}
}

func (c *Counter) Listen() {
	counter := 0
	for {
		if counter >= c.limit {
			if err := c.done(); err != nil {
				c.Err <- err
			}
		}
		select {
		case <-c.ctx.Done():
			return
		case <-c.C:
			counter++
		}
	}
}

func (c *Counter) Close() {
	c.cancel()
}

func NewCounter(limit int, done func() error) *Counter {
	ctx, cancel := context.WithCancel(context.Background())
	return &Counter{
		ctx:    ctx,
		cancel: cancel,
		limit:  limit,
		done:   done,
		C:      make(chan struct{}, limit),
		Err:    make(chan error, limit),
	}
}

type Client interface {
	GetID() (int64, error)
	Recycle(id int64) error
	Close()
}

type pool struct {
	// DON'T BREAK FIELD ORDER
	size        int64
	recycleSize int64

	pool        chan int64
	recyclePool chan int64

	*Dialer

	c *Counter

	conn *grpc.ClientConn
}

func (p *pool) Recycle(id int64) error {
	if atomic.LoadInt64(&p.recycleSize) >= int64(p.MaxPoolSize*4) {
		return errors.New("recyclePool is full")
	}
	atomic.AddInt64(&p.recycleSize, 1)
	p.recyclePool <- id
	return nil
}

func (p *pool) GetID() (int64, error) {
	// the recycle pool takes precedence
	if atomic.LoadInt64(&p.recycleSize) > 0 {
		atomic.AddInt64(&p.recycleSize, -1)
		id := <-p.recyclePool
		return id, nil
	}

	if atomic.LoadInt64(&p.size) == 0 {
		return 0, errors.New("get next id failed, id pool is empty")
	}
	atomic.AddInt64(&p.size, -1)
	id := <-p.pool
	p.c.Call()
	return id, nil
}

func (p *pool) BulkFetchID() error {
	bulkSize := int64(p.MaxPoolSize)
	remSize := atomic.LoadInt64(&p.size)
	if remSize <= remSize {
		bulkSize = bulkSize - remSize
	}
	ctx, cancel := context.WithCancel(p.ctx)
	defer cancel()
	reqId, _ := rand.Int(rand.Reader, big.NewInt(MaxRequestId))
	reply, err := idService.NewOceanIDClient(p.conn).BulkGenerateID(ctx, &idService.IDBulkRequest{
		DC:        p.DcID,
		WorkerId:  p.WorkerId,
		RequestId: reqId.Uint64(),
		BulkSize:  uint32(bulkSize),
	})
	if err != nil {
		return errors.Wrap(err, "request to assign ID failed")
	}
	if reply.ReplyId != reqId.Uint64() {
		return errors.New("invalid ID reply, reply_id verification failed")
	}
	for _, id := range reply.Ids {
		p.pool <- id
		atomic.AddInt64(&p.size, 1)
	}
	return nil
}

func (p *pool) Close() {
	p.c.Close()
}
