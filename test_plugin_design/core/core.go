package core

import (
	"context"
	"log"

	"go_lab/test_plugin_design/broker"
)


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
