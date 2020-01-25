package core

import (
	"go_lab/test_plugin_design/broker"
	"context"
	"log"
)

type Options struct {
	Broker  broker.Broker
	Context context.Context
}

type Option func(*Options)

// Broker 插件拔插
func Broker(b broker.Broker, bOpts ...broker.Option) Option {
	return func(o *Options) {
		o.Broker = b
		log.Println("--- replace Broker instance")
		b.Init(bOpts...)
	}
}
