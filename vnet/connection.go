package vnet

import (
	"context"
	"net"
)

// Connection tcp 连接
type Connection struct {
	Conn     *net.Conn    // 客户端conn
	TcpConn  *net.TCPConn // 服务端conn
	Ctx      context.Context
	Cancel   context.CancelFunc
	ErrTimes int // 异常次数
}
