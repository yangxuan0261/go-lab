package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
)

func test_002() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	data := []byte(`{
		"Hello": {
			"Name": "leaf"
		}
	}`)

	// len + data
	m := make([]byte, 2+len(data))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(len(data)))

	copy(m[2:], data)

	// 发送消息
	conn.Write(m)

	// close
	var wg sync.WaitGroup
	c := make(chan os.Signal, 1)

	wg.Add(1)
	go readMsg11(&wg, c, conn)
	wg.Wait()

	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	// conn.Close() // 客户端调了这个关闭, 服务器 read message: EOF
	fmt.Printf("Leaf closing down (signal: %v)\n", sig)
}

func readMsg11(wg *sync.WaitGroup, c chan os.Signal, conn net.Conn) {
	wg.Done()

	for {
		//  var buf [50]byte
		buf := make([]byte, 2)
		fmt.Println("--- readMsg")
		_, err := conn.Read(buf) // 这个是阻塞的
		if err != nil {
			fmt.Println("conn closed")
			c <- os.Kill
			return
		} else {
			dataLen := binary.BigEndian.Uint16(buf) // 读长度
			if dataLen > 0 {
				buf = make([]byte, dataLen)
				n, err2 := conn.Read(buf)
				if err2 != nil {
					fmt.Println("conn closed")
					c <- os.Kill
					return
				} else {
					fmt.Println("recv msg:", string(buf[:n]))
				}
			}
		}
	}
}
