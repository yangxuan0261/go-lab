package grpcbroker

import (
	"GoLab/test_plugin_design/broker"
	"log"
)

func NewGrpc(opts ...broker.Option) broker.Broker {
	h := &grpcBroker{
		name: "--- grpcBroker",
	}
	return h
}

// 自定义插件
type grpcBroker struct {
	name string
}

func (h *grpcBroker) Init(opts ...broker.Option) error {
	log.Println("--- grpcBroker.Init, name:", h.name)
	return nil
}
