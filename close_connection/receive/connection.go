package receive

import (
	"fmt"
	"io"
	"net"

	"github.com/gobwas/ws"
)

func Connection(conn net.Conn) {
	defer conn.Close()

	for {
		header, err := ws.ReadHeader(conn)
		if err != nil {
			fmt.Println("can not read header", err)
		}

		payload := make([]byte, header.Length)
		_, err = io.ReadFull(conn, payload)
		if err != nil {
			fmt.Println("can not read payload", err)
		}
		if header.Masked {
			ws.Cipher(payload, header.Mask, 0)
		}

		// Reset the Masked flag, server frames must not be masked as
		// RFC6455 says.
		header.Masked = false

		if err := ws.WriteHeader(conn, header); err != nil {
			fmt.Println("can not write header", err)
		}

		if err = doSomethingWithData(payload, conn); err != nil {
			fmt.Println("can not handle data", err)
		}

		if header.OpCode == ws.OpClose {
			return
		}
	}
}

func doSomethingWithData(data []byte, conn net.Conn) error {
	fmt.Println(string(data))
	_, err := conn.Write(data)
	return err
}
