package main

import (
	proto2 "go_lab/test_net/test_socket/proto"
	"go_lab/test_net/test_socket/test_tcp"
	"bufio"
	"context"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"os"
	"sync"

	"time"

	"github.com/golang/protobuf/proto"
)

func send(info *CInfo) {
	info.wg.Add(1)
	defer info.wg.Done()

	cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
		if sender.Text() == "stop" {
			info.conn.Close()
			return
		}

		if info.conn == nil {
			return
		}

		cnt++
		stSend := &proto2.UserInfo{
			Message: sender.Text(),
			Length:  *proto.Int(len(sender.Text())),
			Cnt:     *proto.Int(cnt),
		}

		//protobuf编码
		pData, err := proto.Marshal(stSend)
		if err != nil {
			panic(err)
		}

		//发送

		err = test_tcp.WriteBuff(info.conn, pData)
		if err != nil {
			info.conn.Close()
			return
		}
	}
}

func recv(info *CInfo) {
	info.wg.Add(1)
	defer info.wg.Done()

	info.wg.Done() // 抵消 main 中的 Add

	for {
		pData, err := test_tcp.ReadBuff(info.conn)
		if err != nil { // 客户端主动断线
			log.Println("--- conn.Read, err:", err)
			info.conn.Close()
			info.conn = nil
			return
		}

		stReceive := &proto2.UserInfo{}
		err = proto.Unmarshal(pData, stReceive) //protobuf 解码
		if err != nil {
			log.Println("--- proto.Unmarshal, err:", err)
			return
		}

		log.Println("receive", info.conn.RemoteAddr(), stReceive)

	}
}

type CInfo struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
	conn   net.Conn
}

func getTls() *tls.Config {
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
		RootCAs:            certPool,
		InsecureSkipVerify: true,
	}
	return tlsCfg
}

func main() {
	addr := "localhost:6600"
	var conn net.Conn
	var err error

	//for conn, err = tls.Dial("tcp", addr, getTls()); err != nil; conn, err = net.Dial("tcp", addr) {
	for conn, err = net.Dial("tcp", addr); err != nil; conn, err = net.Dial("tcp", addr) {
		log.Printf("--- connect addr:%s fail\n", addr)
		time.Sleep(time.Second)
		log.Println("reconnect...")
	}
	log.Printf("--- connect addr:%s success\n", addr)
	defer conn.Close()

	var wg sync.WaitGroup
	info := &CInfo{
		wg:   &wg, // struct, 不允许复制拷贝, 只能用指针的形式传递
		conn: conn,
	}

	info.wg.Add(1)

	go send(info)
	go recv(info)

	info.wg.Wait()
	log.Println("--- exit main")
}
