package test

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"net"
	"testing"
	"time"
	"vgo/lib"
	_ "vgo/vlog"
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

	for {
		time.Sleep(1 * time.Second)
		buff := bytes.NewBuffer([]byte{})

		data := "你好哇!~"

		msg := &lib.Message{Len: uint32(len(data)), Data: []byte(data)}
		binary.Write(buff, binary.LittleEndian, msg.Len)
		binary.Write(buff, binary.LittleEndian, msg.Data)

		conn.Write(buff.Bytes())
	}
}
