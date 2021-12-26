package test

import (
	"testing"
	"time"
	"vgo/vmqtt"
)

func Test_MqttClient(t *testing.T) {
	client := &vmqtt.Client{
		ClientId:  "client_test",
		KeepAlive: 60 * time.Second,
	}
	client.Connect(&vmqtt.Broker{
		Schema: "mqtts",
		Host:   "broker-cn.emqx.io",
		Port:   8883,
	})

	defer client.Mqtt.Disconnect(3)

	client.Mqtt.Subscribe("xx", 0, client.MsgHandler)

	for {
		time.Sleep(2 * time.Second)
		msg := time.Now().Format("2006-01-02 15:04:05")
		client.Publish("xx", 0, false, []byte(msg))
	}

}
