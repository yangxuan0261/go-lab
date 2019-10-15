package main

import (
	stProto "GoLab/test_socket/proto"
	"bufio"
	"log"
	"net"
	"os"

	"time"

	"github.com/golang/protobuf/proto"
)

func main() {
	addr := "localhost:6600"
	var conn net.Conn
	var err error

	//连接服务器
	for conn, err = net.Dial("tcp", addr); err != nil; conn, err = net.Dial("tcp", addr) {
		log.Printf("--- connect addr:%s fail\n", addr)
		time.Sleep(time.Second)
		log.Println("reconnect...")
	}
	log.Printf("--- connect addr:%s success\n", addr)
	defer conn.Close()

	//发送消息
	cnt := 0
	sender := bufio.NewScanner(os.Stdin)
	for sender.Scan() {
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
		conn.Write(pData)
		if sender.Text() == "stop" {
			return
		}
	}
}
