package main

import (
	"GoLab/test_mqtt/work"
	"encoding/json"
	"fmt"

	goprotobuf "GoLab/test_protobuf/proto"

	MQTT "github.com/eclipse/paho.mqtt.golang"
	"github.com/golang/protobuf/proto"

	"sync"
	"time"
)

type ResBody struct {
	Error  string
	Result []byte
}

type ProtoSend struct {
	topic string
	buff  []byte
}

func tcp_protobuf(wg *sync.WaitGroup, ps chan *ProtoSend) {
	defer wg.Done()

	sig := make(chan bool, 1)

	this := new(work.MqttWork)
	opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
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

	sendFn := func(topic string, body []byte) {
		if topic == "" {
			sig <- true
			return
		}
		if body == nil {
			sig <- true
			return
		}

		msg, err := this.Request(topic, body)
		if err != nil {
			fmt.Println("返回消息失败, ", err.Error())
			sig <- true
			return
		}

		// 解码 mqtt 消息体
		rb := &ResBody{}
		retBytes := msg.Payload()
		if err = json.Unmarshal([]byte(retBytes), &rb); err != nil {
			fmt.Println("mqtt json 解码失败:", rb)
			sig <- true
			return
		}

		// 解码 protobuf 消息体
		retmsg := &goprotobuf.HelloWorld{}
		if err = proto.Unmarshal(rb.Result, retmsg); err == nil {
			// fmt.Println("len:", len(rb.Result), "bytes:", rb.Result)
			fmt.Println("Error:", rb.Error)
			fmt.Println(fmt.Sprintf("msgId:%d, topic:%s, body:%s", msg.MessageID(), msg.Topic(), retmsg.String()))
		}
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
func test_multiProtobuf(num int, isPing bool) {
	// c := make(chan os.Signal, 1)
	var wg sync.WaitGroup

	createPsFn := func(flag int, sendCnt int) *ProtoSend {
		msg := &goprotobuf.HelloWorld{
			Id:  proto.Int32(996),
			Str: proto.String(fmt.Sprintf("wilker-%d , call cnt:%d", flag, sendCnt)),
		}

		buffer, err := proto.Marshal(msg)
		if err != nil {
			return nil
		}

		return &ProtoSend{
			topic: "HelloWorld@HelloWorld001/HD_Walk",
			buff:  buffer,
		}
	}

	createTcpFn := func(flag int) {
		ps := make(chan *ProtoSend, 1)
		go tcp_protobuf(&wg, ps)

		cnt := 1
		for {
			ctorPs := createPsFn(flag, cnt)
			ps <- ctorPs
			if !isPing || ctorPs == nil {
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

/*
func test_protobuf() {
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

	sendFn := func(topic string, body []byte) {
		if topic == "" {
			topic = "HelloWorld@HelloWorld001/HD_Walk"
		}
		if body == nil {
			body = []byte("Sorry msg")
		}
		msg, err := this.Request(topic, body)
		if err != nil {
			fmt.Println(err.Error())
		}

		retmsg := &goprotobuf.HelloWorld{}
		retBytes := msg.Payload()

		rb := &ResBody{}

		if err = json.Unmarshal([]byte(retBytes), &rb); err != nil {
			fmt.Println("mqtt json 解码失败:", rb)
		}

		if err = proto.Unmarshal(rb.Result, retmsg); err == nil {
			fmt.Println("len:", len(rb.Result), "bytes:", rb.Result)
			fmt.Println(fmt.Sprintf("msgId:%d, topic:%s, body:%s", msg.MessageID(), msg.Topic(), retmsg.String()))
		}
	}

	msg := &goprotobuf.HelloWorld{
		Id:  proto.Int32(996),
		Str: proto.String("what the fuck"),
	}

	if buffer, err := proto.Marshal(msg); err == nil {
		sendFn("HelloWorld@HelloWorld001/HD_Walk", buffer)
	}

	select {}
}
*/
