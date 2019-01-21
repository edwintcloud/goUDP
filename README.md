# goUDP
This is a project for learning purposes writing a simple UDP server and client in Go. Resources used: godoc.org, [this](https://ops.tips/blog/udp-client-and-server-in-go/), and [this](https://varshneyabhi.wordpress.com/2014/12/23/simple-udp-clientserver-in-golang/) blog post.

## Getting Started
### Server
1. Start the server: `go run server.go`

### Client
1. Start the client: `go run client.go`
2. Send a message by typing in the console.
3. Check the server and see that the packet was received.

## TODO
1. ~~Server~~
2. ~~Client~~
3. Docker Compose for server and client (alpine linux, compiled binary, separate ip addresses)
4. ~~Update README with Getting Started section~~
5. Server - Go Routine, Context with Deadline
6. Server - Spawn hundreds of Go Routines listening on a range of ports
