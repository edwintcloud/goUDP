package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	serverAddr = "127.0.0.1:10001"
	clientAddr = "127.0.0.1:0"
)

// Client is our struct to hold the udp connection
type Client struct {
	C *net.UDPConn
}

func main() {
	client := Client{}

	// connect to server
	err := client.ConnectToServer()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Printf("Connected to server at: %s", serverAddr)
	}

	// defer connection to close when main go routine exits
	defer client.C.Close()

	// prompt for messages and send each message to server
	for {
		msg := ""
		fmt.Print("\nEnter a message to send: ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			msg = scanner.Text()
		}
		err = client.SendMessage(msg)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Message Sent!")
		}
	}

}

// ConnectToServer connects to udp server
func (c *Client) ConnectToServer() error {

	// resolve serverAddr to fully qualified address
	server, err := net.ResolveUDPAddr("udp", serverAddr)
	if err != nil {
		return err
	}

	// resolve clientAddr to fully qualified address
	client, err := net.ResolveUDPAddr("udp", clientAddr)
	if err != nil {
		return err
	}

	// connect to udp server
	c.C, err = net.DialUDP("udp", client, server)
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil

}

// SendMessage sends a string message to server
func (c *Client) SendMessage(msg string) error {

	// send message to server
	_, err := c.C.Write([]byte(msg))
	if err != nil {
		return err
	}

	// if all went well, return nil
	return nil

}
