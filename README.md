# websocket-example
A repository to test a websocket implementation based on Nerzal's talk about websockets in Go.

## What is a websocket?
You may know http protocol, which is basically: client sends requests and server sends response. It's an unidirectional protocol (client -> server). WebSocket is a bidirectional protocol (client <-> server). It is used for exchanging data in real time. The connection between client and server is being kept alive until one of both parties decides to close it.

## First steps 
First we want to try the example given by Nerzal. All of the code concerning this section can be found in the directory `first_steps`

### Listen for connection
In the first step we need to establish a biderictional connection. This is a connection, to which clients can connect later on. Our strategy is to listen for announces, accept a connection and then upgrade it to a websocket. 
- We start to listen in `first_steps/init.go`. A main function starts a *Listener*, which basically listens for any announces on a certain port. 
- It can accept a connection with the `Accept` function.
- After a connection exists, an Upgrader object is created, which upgrades the connection to a WebSocket.

### Handle data
In the second step, when a connection is established, we need to handle incoming (client) data. The data comes in so-called frames, which is a transmission unit. This is done in `receive/connection.go` . 
- We create a reader, which is supposed to read frames from a connection with a server
- It reads a frame, reads all the bytes from a frame and finally handles the previously read byte array

### Send messages
- Run the code with `go run init.go`
- Client: You can install a web socket test client (chrome extension) to test your program. Paste the url ws://127.0.0.1:1337/ and click "Open" to establish connection.
- Server: After opening the client you will receive some messages, which will be displayed in your client

## Closing the connection from client side
In the previous example I experienced a problem while closing the connection. The server can close the connection, but if a client tries to close the connection, an error is thrown. This needs to be fixed. It is necessary to dive deeper into _gobwas_ lib _https://github.com/gobwas/ws_ which Nerzal mentioned in his talk.

All of the code concerning this section can be found in the directory `closing_connection`

### Data Frames
To understand how to handle connections gracefully, we need to dig a little deeper into dataframes. Broadly speaking, data frames consist of a header and payload. The following information from the header are interesting for us:
- **Payload length**: To read the payload data, we need to know when to stop reading
- **Opcode**: Defines how to interpret the payload data (important, when a connection is closed)
- **Mask**: The specification of WebSocket protocol defines, that messages from client to server have to be encoded, while messages from server to client must not be encoded