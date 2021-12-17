package test

import (
	"testing"
	_ "vgo/vlog"
	"vgo/vnet"
)

func Test_TcpClient(t *testing.T) {
	c := &vnet.Client{}
	c.Start("127.0.0.1", 9966)
	select {}
}

func Test_TcpServer(t *testing.T) {
	s := &vnet.Server{Host: "127.0.0.1", Port: 9966}
	s.Start()

	select {}
}
