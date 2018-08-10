package main

import mqtt "github.com/eclipse/paho.mqtt.golang"

func main() {
	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.hivemq.com:1883").SetClientID("sample")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
}
