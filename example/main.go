package main

import (
	"log"

	"github.com/singzer/hrp01"
)

func main() {
	server := hrp01.New("0.0.0.0:9999")

	server.OnNewClient(func(c *hrp01.Client) {
		// new client connected
		// lets send some message
		log.Println("new client connected")
		c.Send("Hello")
	})
	server.OnNewMessage(func(c *hrp01.Client, message string) {

		log.Println("message", message)
		// new message received
	})
	server.OnClientConnectionClosed(func(c *hrp01.Client, err error) {
		log.Println("client connection closed")
		// connection with client lost
	})

	server.Listen()
}
