package broker

import (
	"context"
	"log"
)

type Options struct {
	Addrs   string
	Context context.Context
}

type Option func(*Options)

type Broker interface {
	Init(...Option) error
}

var (
	DefaultBroker Broker = newHttpBroker()
)

func newHttpBroker(opts ...Option) Broker {
	h := &httpBroker{
		name: "--- httpBroker",
	}

	return h
}

// 自定义插件, -- 此处为默认插件
type httpBroker struct {
	name string
}

func (h *httpBroker) Init(opts ...Option) error {
	log.Println("--- httpBroker.Init")

	optsIns := &Options{
		Addrs: "--- httpBroker addr",
	}

	for _, o := range opts {
		o(optsIns)
	}

	return nil
}

func Addrs(addr string) Option {
	return func(o *Options) {
		o.Addrs = addr
	}
}
