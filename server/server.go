package main

import (
	"fmt"
	"log"
	"net"
)

const (
	port = ":10001"
	bs   = 1 << 10 // 1024 bytes
)

func main() {

	// start udp server
	StartServer()

}

// StartServer starts our udp server and listens for packets
func StartServer() {

	// resolve udp address from port
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		log.Fatal(err)
	}

	// start udp server on addr
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Server listening for packets on %s\n\n", conn.LocalAddr())
	}

	// defer server to close when method finishes
	defer conn.Close()

	// allocate buffer to hold incoming packet data
	buf := make([]byte, bs)

	// listen for packets
	for {

		// blocking - we wait for a packet
		n, sender, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println(err)
		}

		// log the packet
		log.Printf("packet received from: %s, bytes=%d\n", sender.String(), n)

		// Print the data from packet
		fmt.Printf("Message: %s\n\n", string(buf[0:n]))

	}

}
