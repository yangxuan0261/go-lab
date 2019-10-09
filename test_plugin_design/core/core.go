package core

import (
	"context"
	"log"

	"GoLab/test_plugin_design/broker"
)

type Options struct {
	Broker  broker.Broker
	Context context.Context
}

type Option func(*Options)

type Service interface {
	Init(...Option)
}

type srv struct {
	opts Options
}

func NewService(opts ...Option) Service {
	log.Println("--- NewService")
	opt := Options{
		Broker:  broker.DefaultBroker,
		Context: context.Background(),
	}

	for _, o := range opts {
		o(&opt)
	}

	return &srv{
		opts: opt,
	}
}

func (s *srv) Init(opts ...Option) {
	log.Println("--- Init")

}

// Broker 插件拔插
func Broker(b broker.Broker, bOpts ...broker.Option) Option {
	return func(o *Options) {
		o.Broker = b
		log.Println("--- replace Broker instance")
		b.Init(bOpts...)
	}
}
