package main

import (
	_ "vgo/vlog"
	"vgo/vnet"
)

func main() {
	s := &vnet.Server{Host: "127.0.0.1", Port: 9966}

	go s.Start()

	select {}
}
