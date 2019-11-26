package broker

import "context"

type Options struct {
	Addrs   string
	Context context.Context
}

type Option func(*Options)

func Addrs(addr string) Option {
	return func(o *Options) {
		o.Addrs = addr
	}
}