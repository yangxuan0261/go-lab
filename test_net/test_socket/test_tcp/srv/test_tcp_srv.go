package main

import (
	proto2 "GoLab/test_net/test_socket/proto"
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"

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

	cnt, err := this.conn.Read(buf) //读消息
	if err != nil {                 // 客户端主动断线
		log.Println("--- CAgent.ReadMsg, err:", err)
		this.cancel()
		return
	}

	stReceive := &proto2.UserInfo{}
	pData := buf[:cnt]

	err = proto.Unmarshal(pData, stReceive) //protobuf 解码
	if err != nil {
		log.Println("--- proto.Unmarshal, err:", err)
		return
	}

	log.Println("receive", this.conn.RemoteAddr(), stReceive)
	if stReceive.Message == "stop" {
		this.cancel()
	}
}

// TLS认证
func getTlsCfg() *tls.Config {
	cert, err := tls.LoadX509KeyPair("../conf/server.pem", "../conf/server.key")
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile("../conf/ca.pem")
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	tlsCfg := &tls.Config{
		Certificates:       []tls.Certificate{cert},
		ClientAuth:         tls.RequireAndVerifyClientCert,
		ClientCAs:          certPool,
		InsecureSkipVerify: false,
	}
	return tlsCfg
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	//监听
	addr := "localhost:6600"
	// listener, err := net.Listen("tcp", addr)
	listener, err := tls.Listen("tcp", addr, getTlsCfg())
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
	cancel()
	log.Println("--- exist, signal 222:", s)
}
