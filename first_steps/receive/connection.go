package receive

import (
	"io/ioutil"
	"net"

	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// Connection establishes a websocket connection, to which a client can "subscribe".
// Handles messages which are received from the client.
func Connection(conn net.Conn) {
	reader := wsutil.NewReader(conn, ws.StateServerSide)

	// receive message
	for {
		_, err := reader.NextFrame()
		if err != nil {
			// handle error
		}

		data, err := ioutil.ReadAll(reader)
		if err != nil {
			// handle error
		}
		doSomethingWithData(data)
	}
}

func doSomethingWithData(data []byte) {
	println(string(data))
}
