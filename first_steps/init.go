package main

import (
	"net"
	"time"

	"websocket-example/first_steps/receive"
	"websocket-example/first_steps/send"

	"github.com/gobwas/ws"
)

// TODO: handle closing connection
func main() {
	// init
	listener, err := net.Listen("tcp", "127.0.0.1:1337")
	if err != nil {
		println(err.Error())
	}

	conn, err := listener.Accept()
	if err != nil {
		println(err.Error())
	}

	upgrader := ws.Upgrader{}
	if _, err = upgrader.Upgrade(conn); err != nil {
		println(err.Error())
	}

	// example for sending message from server to client
	go func(conn net.Conn) {
		for i := 0; i < 5; i++ {
			send.Message(conn, "Hello client")
			time.Sleep(1 * time.Second)
		}
	}(conn)

	receive.Connection(conn)
}
