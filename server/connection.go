package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type connection struct {
	ws     *websocket.Conn // The websocket connection.
	send   chan []byte     // Buffered channel of outbound messages.
	client *Client         // The client attached to this connection.
	h      *hub            // The hub this connection is talking to.
}

func (c *connection) reader() {
	for {
		_, message, err := c.ws.ReadMessage()
		if err != nil {
			break
		}
		fmt.Print("message received")
		fmt.Print(message)
		//c.h.broadcast <- message
	}
	c.ws.Close()
}

func (c *connection) writer() {
	for message := range c.send {
		err := c.ws.WriteMessage(websocket.BinaryMessage, message)
		if err != nil {
			break
		}
	}
	c.ws.Close()
}
