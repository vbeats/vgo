package vnet

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"strconv"
	"time"
	_ "vgo/vlog"
)

// Server tcp server
type Server struct {
	Host string
	Port int
}

func (s *Server) Start() {
	server, err := net.Listen("tcp", s.Host+":"+strconv.Itoa(s.Port))
	if err != nil {
		logrus.Error("server 启动tcp监听失败...", err)
		panic(err)
	}

	logrus.Info("server start... listen on port: ", s.Port)

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go handleClientConn(&Connection{Conn: &conn, ErrTimes: 0})
	}
}

// 处理连接
func handleClientConn(conn *Connection) {
	logrus.Info("客户端建立连接...", (*conn.Conn).RemoteAddr())

	conn.Ctx, conn.Cancel = context.WithCancel(context.Background())

	defer (*conn.Conn).Close()
	defer conn.Cancel()

	for {
		select {
		case <-conn.Ctx.Done(): // 关闭连接
			logrus.Info("客户端断开连接....", (*conn.Conn).RemoteAddr())
			return
		default:
			time.Sleep(1 * time.Second)
			data := "pong" + time.Now().Format("2006-01-02 15:04:05")
			// 先发后读
			conn.sendClientMsg(Msg{Len: uint32(len(data)), Data: []byte(data)})
			conn.readClientMsg(&Msg{})
		}
	}
}

// 发送msg数据到客户端
func (c *Connection) sendClientMsg(msg Msg) {
	select {
	case <-c.Ctx.Done():
		return
	default:
		buff := bytes.NewBuffer([]byte{})

		binary.Write(buff, binary.LittleEndian, msg.Len)
		binary.Write(buff, binary.LittleEndian, msg.Data)

		(*c.Conn).SetWriteDeadline(time.Now().Add(3 * time.Second))
		_, err := (*c.Conn).Write(buff.Bytes())

		if err != nil {
			logrus.Error("向客户端写数据异常", (*c.Conn).RemoteAddr(), err)
			c.ErrTimes += 1
			if c.ErrTimes > 3 { // 写超时超过3次断开tcp连接
				c.Cancel()
				return
			}
		}
	}
}

// 接收客户端数据到msg中
func (c *Connection) readClientMsg(msg *Msg) {
	select {
	case <-c.Ctx.Done():
		return
	default:
		buff := make([]byte, 4)
		n, err := io.ReadFull(*c.Conn, buff)
		if err != nil {
			logrus.Error("接收客户端数据异常...", err)
			return
		}

		if n > 0 {
			// 消息头长度
			binary.Read(bytes.NewReader(buff[:n]), binary.LittleEndian, &msg.Len)
			// 消息体
			data := make([]byte, msg.Len)
			io.ReadFull(*c.Conn, data)

			msg.Data = data

			logrus.Info("收到客户端消息: ", string(msg.Data))
		}
	}
}
