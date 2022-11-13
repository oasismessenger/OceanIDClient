package IDClient

var (
	DefaultDialerConfig = Dialer{
		MaxPoolSize: 5000,
		MinPoolSize: 2000,
		DcID:        0,
		WorkerId:    1,
	}
)

func NewDialerConfigWithDefault(addr string) *Dialer {
	cfg := DefaultDialerConfig
	cfg.Addr = addr
	return &cfg
}
