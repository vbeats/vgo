package test

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"testing"
	"time"
	_ "vgo/vlog"
	"vgo/vnet"
)

func Test_Tcp(t *testing.T) {
	client, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9966")
	conn, err := net.DialTCP("tcp", nil, client)
	if err != nil {
		logrus.Error("客户端连接失败...", err)
		panic(err)
	}

	defer conn.Close()

	logrus.Info("connected.....!")

	// 写数据
	go func() {
		for {
			time.Sleep(1 * time.Second)
			buff := bytes.NewBuffer([]byte{})

			data := "我是client啊!~" + time.Now().Format("2006-01-02 15:04:05")

			msg := &vnet.Msg{Len: uint32(len(data)), Data: []byte(data)}
			binary.Write(buff, binary.LittleEndian, msg.Len)
			binary.Write(buff, binary.LittleEndian, msg.Data)

			conn.Write(buff.Bytes())
		}
	}()

	// 读数据
	go func() {
		for {
			buff := make([]byte, 4)
			n, err := io.ReadFull(conn, buff)
			if err != nil {
				break
			}

			msg := &vnet.Msg{}

			if n > 0 {
				// 消息头长度
				binary.Read(bytes.NewReader(buff[:n]), binary.LittleEndian, &msg.Len)
				// 消息体
				data := make([]byte, msg.Len)
				io.ReadFull(conn, data)

				msg.Data = data

				logrus.Info("收到服务端消息: ", string(msg.Data), " 消息长度: ", msg.Len)
			}
		}
	}()

	select {}
}
