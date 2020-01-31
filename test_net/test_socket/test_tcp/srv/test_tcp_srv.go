package main

import (
	"go-lab/lib"
	proto2 "go-lab/test_net/test_socket/proto"
	"go-lab/test_net/test_socket/test_tcp"
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
	"net"
)

// 参考: https://segmentfault.com/a/1190000009277748

type CAgent struct {
	uid    uint64
	conn   net.Conn
	ctx    context.Context
	cancel context.CancelFunc
}

func (a *CAgent) Run() {
	for {
		select {
		case <-a.ctx.Done():
			a.Done()
			return
		default:
			a.ReadMsg()
		}
	}
}

func (a *CAgent) Done() {
	log.Printf("--- CAgent.Done, %d exit\n", a.uid)
	a.conn.Close()
}

func (a *CAgent) ReadMsg() {
	pData, err := test_tcp.ReadBuff(a.conn)
	if err != nil { // 客户端主动断线
		log.Println("--- CAgent.ReadMsg, err:", err)
		a.cancel()
		return
	}

	stReceive := &proto2.UserInfo{}

	err = proto.Unmarshal(pData, stReceive) //protobuf 解码
	if err != nil {
		log.Println("--- proto.Unmarshal, err:", err)
		a.cancel()
		return
	}

	log.Printf("--- receive: addr:%s, data:%+v\n", a.conn.RemoteAddr(), stReceive)

	stReceive.Message = "srv pong"
	pData, err = proto.Marshal(stReceive)
	if err != nil {
		log.Println("--- proto.Marshal, err:", err)
		a.cancel()
		return
	}

	err = test_tcp.WriteBuff(a.conn, pData)
	if err != nil {
		log.Println("--- WriteBuff, err:", err)
		a.cancel()
		return
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
	ctx, _ := context.WithCancel(context.Background())

	//监听
	addr := "localhost:6600"
	listener, err := net.Listen("tcp", addr)
	//listener, err := tls.Listen("tcp", addr, getTlsCfg())
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
	lib.WaitSignal()
	log.Println("--- exist, signal 222:")
}
