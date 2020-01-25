package main

import (
	"go-lab/test_plugin_design/broker"
	"go-lab/test_plugin_design/core"
	"log"
	"go-lab/test_plugin_design/plugins/broker/grpc"
)

// package test_base

func main() {
	log.Println("--- start")

	opt := broker.Addrs("-- hello world")
	gb := grpcbroker.NewGrpc(opt)

	opts := core.NewService(
		core.Broker(gb, opt),
	)
	_ = opts
	log.Println("--- end")
}
