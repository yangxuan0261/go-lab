package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"sync"
	"time"
)

func main() {
	// test_multiJson(3000, true)
	ArgsParser()
}

func ArgsParser() {
	num := flag.Int("num", 1, "arg is num")
	isPing := flag.Bool("isPing", true, "arg is isPing")
	flag.Parse()
	test_multiJson(*num, *isPing)
}

func tcp_json(wg *sync.WaitGroup, ps chan *JsonSend) {
	defer wg.Done()

	sig := make(chan bool, 1)

	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	go readMsg22(sig, conn)

	sendFn := func(data string) {
		m := make([]byte, 2+len(data))
		binary.BigEndian.PutUint16(m, uint16(len(data)))
		copy(m[2:], data)
		conn.Write(m)
	}

	for {
		select {
		case sendBody := <-ps:
			if sendBody != nil {
				sendFn(sendBody.buff)
			} else {
				sig <- true
			}
		case <-sig:
			fmt.Printf("结束发送:(signal: %v)\n", sig)
			return
		}
	}

}

type JsonSend struct {
	buff string
}

// 多开客户端数量, 是否定时发送
func test_multiJson(num int, isPing bool) {
	// c := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	createJsFn := func(flag int, sendCnt int) *JsonSend {
		content := `{
			"Hello": {
				"Name": "leaf, flag-%d, cnt-%d"
			}
		}`
		return &JsonSend{
			buff: fmt.Sprintf(content, flag, sendCnt),
		}
	}

	createTcpFn := func(flag int) {
		js := make(chan *JsonSend, 1)
		go tcp_json(&wg, js)

		cnt := 1
		for {
			ctorJs := createJsFn(flag, cnt)
			js <- ctorJs
			if !isPing || ctorJs == nil {
				break
			}
			cnt++
			time.Sleep(time.Second * 3)
		}
	}

	for index := 1; index <= num; index++ {
		wg.Add(1)
		go createTcpFn(index)
	}

	wg.Wait()
	fmt.Println("GameOver!!!")
}

func readMsg22(c chan<- bool, conn net.Conn) {
	for {
		//  var buf [50]byte
		buf := make([]byte, 2)
		_, err := conn.Read(buf) // 这个是阻塞的
		if err != nil {
			fmt.Println("conn closed")
			c <- true
			return
		} else {
			dataLen := binary.BigEndian.Uint16(buf) // 读长度
			if dataLen > 0 {
				buf = make([]byte, dataLen)
				n, err2 := conn.Read(buf)
				if err2 != nil {
					fmt.Println("conn closed")
					c <- true
					return
				} else {
					fmt.Println("recv msg:", string(buf[:n]))
				}
			}
		}
	}
}
