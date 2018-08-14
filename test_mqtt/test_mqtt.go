package main

import (
	"GoLab/test_mqtt/work"
	"fmt"

	MQTT "github.com/eclipse/paho.mqtt.golang"

	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

func main() {
	doConn(nil, "-1", false)
	// doMultiConn()
}

func doMultiConn() {
	var wg sync.WaitGroup
	cnt := 10000
	for index := 1; index <= cnt; index++ {
		wg.Add(1)
		go doConn(&wg, strconv.Itoa(index), true)
	}

	wg.Wait()
	fmt.Println("程序结束 666")
}

func doConn(wg *sync.WaitGroup, flag string, isPing bool) {
	c := make(chan os.Signal, 1)

	this := new(work.MqttWork)
	opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("连接断开", err.Error())
		c <- os.Kill
	})
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("连接成功")
	})
	err := this.Connect(opts)
	if err != nil {
		fmt.Println(err.Error())
	}

	cnt := 1
	sendFn := func() {
		//访问HelloWorld001模块的HD_Say函数
		hiStr := fmt.Sprintf(`{"say":"我是wilker%s, cnt:%d"}`, flag, cnt)
		msg, err := this.Request("HelloWorld@HelloWorld001/HD_Say", []byte(hiStr))
		if err != nil {
			fmt.Println(err.Error())
		}
		fmt.Println(fmt.Sprintf("topic :%s  body :%s", msg.Topic(), string(msg.Payload())))
	}
	sendFn()

	if isPing {
		for {
			cnt++
			sendFn()
			time.Sleep(time.Second * 3)
		}
	}

	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("mqant closing down (signal: %v)\n", sig)
	fmt.Println("--- flag done:", flag)
	if wg != nil {
		defer wg.Done()
	}
}
