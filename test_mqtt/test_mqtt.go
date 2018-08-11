package main

import (
	"GoLab/test_mqtt/work"
	"fmt"
	"os"
	"os/signal"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

func main() {
	this := new(work.MqttWork)
	opts := this.GetDefaultOptions("tcp://127.0.0.1:3563")
	opts.SetConnectionLostHandler(func(client MQTT.Client, err error) {
		fmt.Println("连接断开", err.Error())
	})
	opts.SetOnConnectHandler(func(client MQTT.Client) {
		fmt.Println("连接成功")
	})
	err := this.Connect(opts)
	if err != nil {
		fmt.Println(err.Error())
	}

	//访问HelloWorld001模块的HD_Say函数
	msg, err := this.Request("HelloWorld@HelloWorld001/HD_Say", []byte(`{"say":"我是梁大帅"}`))
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(fmt.Sprintf("topic :%s  body :%s", msg.Topic(), string(msg.Payload())))

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	sig := <-c
	fmt.Printf("mqant closing down (signal: %v)\n", sig)

}
