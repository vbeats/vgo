package main

import (
	"github.com/sirupsen/logrus"
	"net"
	"vgo/service"
	_ "vgo/vlog"
)

func main() {
	server, err := net.Listen("tcp", ":9966")
	if err != nil {
		logrus.Error("server 启动tcp监听失败...", err)
		panic(err)
	}
	defer server.Close()

	logrus.Info("server 启动tcp监听成功...")

	for {
		conn, err := server.Accept()
		if err != nil {
			continue
		}
		go service.HandleConn(&conn)
	}
}
