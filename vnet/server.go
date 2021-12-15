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
	Conn net.Conn
	ctx  context.Context
}

func (s *Server) Start() {
	server, err := net.Listen("tcp", s.Host+":"+strconv.Itoa(s.Port))
	if err != nil {
		logrus.Error("server 启动tcp监听失败...", err)
		panic(err)
	}

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		s.Conn = conn
		go handleConn(s)
	}
}

// 处理连接
func handleConn(s *Server) {
	logrus.Info("客户端建立连接...", s.Conn.RemoteAddr())
	ctx, cancel := context.WithCancel(context.Background())
	s.ctx = ctx

	defer s.Conn.Close()
	defer cancel()

	go read(s)
	go write(s)

	select {}
}

// 发送数据
func write(s *Server) {
	for {
		select {
		case <-s.ctx.Done(): // 连接关闭
			return
		default:
			time.Sleep(1 * time.Second)
			buff := bytes.NewBuffer([]byte{})

			data := "你好哇!~" + time.Now().Format("2006-01-02 15:04:05")

			msg := &Msg{Len: uint32(len(data)), Data: []byte(data)}
			binary.Write(buff, binary.LittleEndian, msg.Len)
			binary.Write(buff, binary.LittleEndian, msg.Data)
			s.Conn.Write(buff.Bytes())
		}
	}
}

// 接收数据
func read(s *Server) {
	for {
		select {
		case <-s.ctx.Done(): // 连接关闭
			return
		default:
			buff := make([]byte, 4)
			n, err := io.ReadFull(s.Conn, buff)
			if err != nil {
				break
			}

			msg := &Msg{}

			if n > 0 {
				// 消息头长度
				binary.Read(bytes.NewReader(buff[:n]), binary.LittleEndian, &msg.Len)
				// 消息体
				data := make([]byte, msg.Len)
				io.ReadFull(s.Conn, data)

				msg.Data = data

				logrus.Info("收到客户端消息: ", string(msg.Data))
			}
		}
	}
}
