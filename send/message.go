package send

import (
	"net"

	"github.com/gobwas/ws/wsutil"
)

// Message sends a message from server to client through a certain connection
func Message(conn net.Conn, message string) {
	msg := []byte(message)
	err := wsutil.WriteServerText(conn, msg)
	if err != nil {
		panic(err)
	}
}
