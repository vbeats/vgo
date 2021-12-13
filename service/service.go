package service

import (
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"vgo/lib"
	_ "vgo/vlog"
)

// HandleConn 处理连接
func HandleConn(conn *net.Conn) {
	defer (*conn).Close()

	logrus.Info("客户端建立连接...", (*conn).RemoteAddr())

	for {
		buff := make([]byte, 4)
		n, err := io.ReadFull(*conn, buff)
		if err != nil {
			break
		}

		msg := &lib.Message{}

		if n > 0 {
			// 消息头长度
			binary.Read(bytes.NewReader(buff[:n]), binary.LittleEndian, &msg.Len)
			// 消息体
			data := make([]byte, msg.Len)
			io.ReadFull(*conn, data)

			msg.Data = data

			logrus.Info("收到客户端消息: ", string(msg.Data))
		}
	}
}
