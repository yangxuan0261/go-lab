package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"

	stProto "GoLab/test_socket/proto"

	"github.com/golang/protobuf/proto"
)

// 参考: https://segmentfault.com/a/1190000009277748

type CAgent struct {
	uid    uint64
	conn   net.Conn
	ctx    context.Context
	cancel context.CancelFunc
}

func (this *CAgent) Run() {
	for {
		select {
		case <-this.ctx.Done():
			this.Done()
			return
		default:
			this.ReadMsg()
		}
	}
}

func (this *CAgent) Done() {
	log.Printf("--- CAgent.Done, %d exit\n", this.uid)
	this.conn.Close()
}

func (this *CAgent) ReadMsg() {
	buf := make([]byte, 4096, 4096)
	//读消息
	cnt, err := this.conn.Read(buf)
	if err != nil {
		panic(err)
	}

	stReceive := &stProto.UserInfo{}
	pData := buf[:cnt]

	//protobuf解码
	err = proto.Unmarshal(pData, stReceive)
	if err != nil {
		panic(err)
	}

	log.Println("receive", this.conn.RemoteAddr(), stReceive)
	if stReceive.Message == "stop" {
		this.cancel()
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//监听
	addr := "localhost:6600"
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	uid := uint64(1)

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				panic(err)
			}
			log.Println("--- Accept, addr:", conn.RemoteAddr())

			aCtx, aCancel := context.WithCancel(ctx)
			agent := &CAgent{
				uid:    uid,
				conn:   conn,
				ctx:    aCtx,
				cancel: aCancel,
			}
			uid++
			go agent.Run()
		}
	}()

	log.Println("--- lintening addr:", addr)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	s := <-c
	log.Println("--- exist, signal 111:", s)
	cancel()
	log.Println("--- exist, signal 222:", s)
}
