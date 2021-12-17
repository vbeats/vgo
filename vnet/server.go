package vnet

import (
	"bytes"
	"context"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"io"
	"net"
	"strconv"
	"sync"
	"time"
	_ "vgo/vlog"
)

var total int32 = 0 // client连接总数

var lock sync.Mutex

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

	logrus.Infof("server start... listen on port: %d", s.Port)

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
	lock.Lock()
	total += 1
	lock.Unlock()
	logrus.Infof("客户端: %s 建立连接 当前连接总数: %d", (*conn.Conn).RemoteAddr(), total)

	conn.Ctx, conn.Cancel = context.WithCancel(context.Background())

	defer (*conn.Conn).Close()
	defer conn.Cancel()

	for {
		select {
		case <-conn.Ctx.Done(): // 关闭连接
			lock.Lock()
			total -= 1
			lock.Unlock()
			logrus.Infof("客户端: %s 断开连接 当前连接总数: %d", (*conn.Conn).RemoteAddr(), total)
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
			logrus.Errorf("向客户端: %s 写数据异常... %s", (*c.Conn).RemoteAddr(), err)
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
			logrus.Errorf("接收客户端: %s 数据异常... %s", (*c.Conn).RemoteAddr(), err)
			return
		}

		if n > 0 {
			// 消息头长度
			binary.Read(bytes.NewReader(buff[:n]), binary.LittleEndian, &msg.Len)
			// 消息体
			data := make([]byte, msg.Len)
			io.ReadFull(*c.Conn, data)

			msg.Data = data

			logrus.Infof("收到客户端: %s 消息: %s", (*c.Conn).RemoteAddr(), string(msg.Data))
		}
	}
}
