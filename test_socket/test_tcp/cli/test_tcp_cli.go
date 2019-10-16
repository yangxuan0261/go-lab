package main

import (
	stProto "GoLab/test_socket/proto"
	"bufio"
	"context"
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
		if info.conn == nil {
			return
		}

		cnt++
		stSend := &stProto.UserInfo{
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
		info.conn.Write(pData)
		if sender.Text() == "stop" {
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
		buf := make([]byte, 1024, 1024)
		cnt, err := info.conn.Read(buf) //读消息
		if err == nil {
			stReceive := &stProto.UserInfo{}
			pData := buf[:cnt]

			err = proto.Unmarshal(pData, stReceive) //protobuf 解码
			if err != nil {
				log.Println("--- proto.Unmarshal, err:", err)
				return
			}

			log.Println("receive", info.conn.RemoteAddr(), stReceive)
		} else {
			log.Println("--- conn.Read, err:", err)
			info.conn.Close()
			info.conn = nil
			return
		}
	}
}

type CInfo struct {
	ctx    context.Context
	cancel context.CancelFunc
	wg     *sync.WaitGroup
	conn   net.Conn
}

func main() {
	addr := "localhost:6600"
	var conn net.Conn
	var err error

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
