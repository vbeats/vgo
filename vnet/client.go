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

// Client tcp client
type Client struct{}

func (c *Client) Start(host string, port int) {

	for {
		client, _ := net.ResolveTCPAddr("tcp", host+":"+strconv.Itoa(port))
		conn, err := net.DialTCP("tcp", nil, client)
		if err != nil {
			logrus.Error("TCP连接失败, 等待5s重试...", err)
			time.Sleep(5 * time.Second)
			continue
		}

		conn.SetKeepAlive(true)

		handleServerConn(&Connection{TcpConn: conn, ErrTimes: 0})
	}
}

// 处理连接
func handleServerConn(conn *Connection) {
	logrus.Infof("client连接上server端 : %s", (*conn.TcpConn).RemoteAddr())

	conn.Ctx, conn.Cancel = context.WithCancel(context.Background())
	defer (*conn.TcpConn).Close()
	defer conn.Cancel()

	for {
		select {
		case <-conn.Ctx.Done(): // 关闭连接
			logrus.Info("断开连接....")
			return
		default:
			time.Sleep(1 * time.Second)
			data := "ping" + time.Now().Format("2006-01-02 15:04:05")
			// 先发后读
			conn.sendServerMsg(Msg{Len: uint32(len(data)), Data: []byte(data)})
			conn.readServerMsg(&Msg{})
		}
	}
}

// 发送msg数据到服务端
func (c *Connection) sendServerMsg(msg Msg) {
	select {
	case <-c.Ctx.Done():
		return
	default:
		buff := bytes.NewBuffer([]byte{})

		binary.Write(buff, binary.LittleEndian, msg.Len)
		binary.Write(buff, binary.LittleEndian, msg.Data)

		(*c.TcpConn).SetWriteDeadline(time.Now().Add(3 * time.Second))
		_, err := (*c.TcpConn).Write(buff.Bytes())

		if err != nil {
			logrus.Error("向服务端写数据异常...", err)
			c.ErrTimes += 1
			if c.ErrTimes > 3 { // 写超时超过3次断开tcp连接
				c.Cancel()
				return
			}
		}
	}
}

// 接收服务端数据到msg中
func (c *Connection) readServerMsg(msg *Msg) {
	select {
	case <-c.Ctx.Done():
		return
	default:
		buff := make([]byte, 4)
		n, err := io.ReadFull(c.TcpConn, buff)
		if err != nil {
			logrus.Error("接收服务端数据异常...", err)
			return
		}

		if n > 0 {
			// 消息头长度
			binary.Read(bytes.NewReader(buff[:n]), binary.LittleEndian, &msg.Len)
			// 消息体
			data := make([]byte, msg.Len)
			io.ReadFull(c.TcpConn, data)

			msg.Data = data

			logrus.Info("收到服务端消息: ", string(msg.Data))
		}
	}
}
