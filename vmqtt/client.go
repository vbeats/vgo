package vmqtt

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

// Broker mqtt broker config
type Broker struct {
	Schema string
	Host   string
	Port   int
}

// Client mqtt client config
type Client struct {
	ClientId  string
	KeepAlive time.Duration
	Mqtt      mqtt.Client
}

// Connect 连接broker
func (client *Client) Connect(broker *Broker) {
	opts := mqtt.NewClientOptions().AddBroker(broker.Schema + "://" + broker.Host + ":" + strconv.Itoa(broker.Port)).SetClientID(client.ClientId)

	opts.SetKeepAlive(client.KeepAlive)
	// 设置消息回调处理函数
	opts.SetDefaultPublishHandler(client.MsgHandler)
	opts.SetPingTimeout(1 * time.Second)

	c := mqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		logrus.Error("mqtt连接失败....", token.Error())
		panic(token.Error())
	}

	client.Mqtt = c
}

func (*Client) MsgHandler(client mqtt.Client, message mqtt.Message) {
	logrus.Infof("收到mqtt broker的消息: %s", string(message.Payload()))
}

// Subscribe 订阅主题
func (client *Client) Subscribe(topic string, qos byte) {
	if token := client.Mqtt.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		logrus.Errorf("mqtt topic: %s 订阅失败... %s", topic, token.Error())
		return
	}
	logrus.Infof("mqtt client subscribe topic %s success", topic)
}

// Publish 发送消息
func (client *Client) Publish(topic string, qos byte, retained bool, payload []byte) {
	logrus.Infof("mqtt client topic: %s send msg %s", topic, string(payload))
	token := client.Mqtt.Publish(topic, qos, retained, payload)
	token.Wait()
}

// UnSubscribe 取消订阅
func (client *Client) UnSubscribe(topic string) {
	if token := client.Mqtt.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		logrus.Errorf("mqtt topic: %s 取消订阅失败... %s", topic, token.Error())
		return
	}
	logrus.Infof("mqtt client unscribe topic: %s success", topic)
}
