package IDClient

import (
	"context"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/oasismessenger/OceanIDClient/schemes/types"
)

type Dialer struct {
	ctx         context.Context
	MaxPoolSize int
	MinPoolSize int
	DcID        types.EnumOceanDC
	WorkerId    uint32
	Addr        string
}

func (d *Dialer) checkValue() error {
	switch {
	case d.MaxPoolSize == 0:
		return errors.New("MaxPoolSize not allowed to be zero")
	case d.MinPoolSize == 0:
		return errors.New("MinPoolSize not allowed to be zero")
	case d.Addr == "":
		return errors.New("Addr not allowed to be empty")
	case d.MaxPoolSize < d.MinPoolSize:
		return errors.New("invalid pool size")
	}
	return nil
}

func (d *Dialer) DialWithContext(ctx context.Context) (Client, error) {

	if err := d.checkValue(); err != nil {
		return nil, err
	}

	conn, err := grpc.DialContext(
		ctx,
		d.Addr,
		grpc.WithDisableRetry(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	d.ctx = ctx

	p := &pool{
		size:        0,
		recycleSize: 0,
		pool:        make(chan int64, d.MaxPoolSize),
		recyclePool: make(chan int64, d.MaxPoolSize*4),
		Dialer:      d,
		conn:        conn,
	}

	p.c = NewCounter(d.MinPoolSize/10, p.BulkFetchID)

	go p.c.Listen()

	return p, nil
}

func (d *Dialer) Dial() (Client, error) {
	return d.DialWithContext(context.Background())
}

func QuickDial(addr string) (Client, error) {
	return NewDialerConfigWithDefault(addr).Dial()
}

func QuickDialWithContext(ctx context.Context, addr string) (Client, error) {
	return NewDialerConfigWithDefault(addr).DialWithContext(ctx)
}
