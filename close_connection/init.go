package main

import (
	"fmt"
	"log"
	"net"
	"websocket-example/close_connection/receive"

	"github.com/gobwas/ws"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:1337")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("can not accept connection")
		}
		_, err = ws.Upgrade(conn)
		if err != nil {
			fmt.Println("can not upgrade to websocket")
		}

		go receive.Connection(conn)
	}
}
