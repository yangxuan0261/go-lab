package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
)

func main() {
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
	go readMsg(&wg, c, conn)

	signal.Notify(c, os.Interrupt, os.Kill)
	wg.Wait()

	// conn.Close() // 客户端调了这个关闭, 服务器 read message: EOF
	sig := "asd" //<-c
	fmt.Printf("Leaf closing down (signal: %v)\n", sig)
}

func readMsg(wg *sync.WaitGroup, c chan os.Signal, conn net.Conn) {
	for {
		//  var buf [50]byte
		buf := make([]byte, 2)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("conn closed")
			c <- os.Kill
		} else {
			dataLen := binary.BigEndian.Uint16(buf) // 读长度
			if dataLen > 0 {
				buf = make([]byte, dataLen)
				n, err2 := conn.Read(buf)
				if err2 != nil {
					fmt.Println("conn closed")
					c <- os.Kill
				} else {
					fmt.Println("recv msg:", string(buf[:n]))
				}
			}
		}

		select {
		case <-c:
			wg.Done()
			fmt.Println("结束读取")
			return
		}
	}
}
