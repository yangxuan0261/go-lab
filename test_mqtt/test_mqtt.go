package main

import (
	"GoLab/test_mqtt/work"
	"flag"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"sync"
	"time"
)

func main() {
	// doConn(nil, "-1", false)
	// test_multiJson(1, false)
	// test_multiProtobuf(1, true)
	// test_multiProtobuf(10, true)
	// test_multiProtobuf(10, true)
	ArgsParser()
}

/*
func doConnWithInput() {
	rInputFn := func(c chan string) {
		for {
			is := ""
			// fmt.Println("Please enter some input: ")
			fmt.Scan(&is)
			c <- is
			// inputReader := bufio.NewReader(os.Stdin)
			// input, err := inputReader.ReadString('\n')
			// if err == nil {
			// 	c <- input
			// }
		}
	}

	dealFn := func(c chan string) {
		this := new(work.MqttWork)
		opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
		opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
			fmt.Println("连接断开", err.Error())
			// c <- os.Kill
		})
		opts.SetOnConnectHandler(func(client MQTT.Client) {
			fmt.Println("连接成功")
		})
		err := this.Connect(opts)
		if err != nil {
			fmt.Println("连接错误:", err.Error())
			// c <- os.Kill
		}

		parseFn := func(src *string) (string, string) {
			sArr := strings.Split(*src, ";")
			if len(sArr) >= 2 {
				return sArr[0], sArr[1]
			}
			return "", ""
		}

		sendFn := func(topic string, body string) {
			if topic == "" {
				topic = "HelloWorld@HelloWorld001/HD_Say"
			}
			if body == "" {
				body = "Sorry msg"
			}
			hiStr := fmt.Sprintf(`{"say":"msg:%s"}`, body)
			msg, err := this.Request(topic, []byte(hiStr))
			if err != nil {
				fmt.Println(err.Error())
			}
			fmt.Println(fmt.Sprintf("topic:%s, body:%s\n", msg.Topic(), string(msg.Payload())))
		}

		for {
			select {
			case msg := <-c:
				// fmt.Println("recv msg:", msg)
				sendFn(parseFn(&msg))
			}
		}
	}

	rc := make(chan string)
	go rInputFn(rc)
	go dealFn(rc)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("mqant closing down (signal: %v)\n", sig)
}
*/

type JsonSend struct {
	topic string
	buff  string
}

func tcp_json(wg *sync.WaitGroup, ps chan *JsonSend) {
	defer wg.Done()

	sig := make(chan bool, 1)

	this := new(work.MqttWork)
	opts := this.GetDefaultOptions("tcp://127.0.0.1:3564")
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("连接断开", err.Error())
		sig <- true
	})
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("连接成功")
	})
	err := this.Connect(opts)
	if err != nil {
		fmt.Println("连接错误:", err.Error())
		return
	}

	// 注册服务端推送
	this.On("Wilker/TestChat", func(client MQTT.Client, pushMsg MQTT.Message) {
		fmt.Println("server push, regMsg:", string(pushMsg.Payload()))
	})

	sendFn := func(topic string, body string) {
		if topic == "" {
			sig <- true
			return
		}

		hiStr := fmt.Sprintf(`{"say":"msg:%s"}`, body)
		msg, err := this.Request(topic, []byte(hiStr))
		if err != nil {
			fmt.Println("返回消息失败, ", err.Error())
			sig <- true
			return
		}

		// 解码 mqtt 消息体
		// rb := &ResBody{}
		// retBytes := msg.Payload()
		// if err = json.Unmarshal([]byte(retBytes), &rb); err != nil {
		// 	fmt.Println("mqtt json 解码失败:", rb)
		// 	sig <- true
		// 	return
		// }

		fmt.Println("返回结果:", string(msg.Payload()))
	}

	for {
		select {
		case sendBody := <-ps:
			if sendBody != nil {
				sendFn(sendBody.topic, sendBody.buff)
			} else {
				sig <- true
			}
		case <-sig:
			fmt.Printf("结束发送:(signal: %v)\n", sig)
			return
		}
	}

}

// 多开客户端数量, 是否定时发送
func test_multiJson(num int, isPing bool) {
	// c := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	createJsFn := func(flag int, sendCnt int) *JsonSend {
		return &JsonSend{
			topic: "HelloWorld@HelloWorld001/HD_Say",
			buff:  fmt.Sprintf("wilker-%d , call cnt:%d", flag, sendCnt),
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

func ArgsParser() {
	num := flag.Int("num", 1, "arg is num")
	isPing := flag.Bool("isPing", true, "arg is isPing")
	isProto := flag.Bool("isProto", true, "arg is isProto")
	flag.Parse()

	if *isProto {
		test_multiProtobuf(*num, *isPing)
	} else {
		test_multiJson(*num, *isPing)
	}
}
