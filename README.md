# websocket-example
A repository to test a websocket implementation based on Nerzal's talk about websockets in Go.

## What is a websocket?
You may know http protocol, which is basically: client sends requests and server sends response. It's an unidirectional protocol (client -> server). WebSocket is a bidirectional protocol (client <-> server). It is used for exchanging data in real time. The connection between client and server is being kept alive until one of both parties decides to close it.

## First steps 
First we want to try the example given by Nerzal.

### Listen for connection
In the first step we need to establish a biderictional connection. This is a connection, to which clients can connect later on. Our strategy is to listen for announces, accept a connection and then upgrade it to a websocket. 
- We start to listen in `init.go`. A main function starts a *Listener*, which basically listens for any announces on a certain port. 
- It can accept a connection with the `Accept` function.
- After a connection exists, an Upgrader object is created, which upgrades the connection to a WebSocket.

### Handle data
In the second step, when a connection is established, we need to handle incoming (client) data. The data comes in so-called frames, which is a transmission unit. This is done in `receive/connection.go` . 
- We create a reader, which is supposed to read frames from a connection with a server
- It reads a frame, reads all the bytes from a frame and finally handles the previously read bytes

### Send messages
- Run the code with `go run init.go`
- Client: You can install a web socket test client (chrome extension) to test your program. Paste the url ws://127.0.0.1:1337/ and click open to establish connection.
- Server: After opening the client you will receive some messages, which will be displayed in your client

