package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type hub struct {
	server *WebServer
	// Registered connections.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection
}

func newHub(server *WebServer) *hub {
	return &hub{
		server:      server,
		broadcast:   make(chan []byte),
		register:    make(chan *connection),
		unregister:  make(chan *connection),
		connections: make(map[*connection]bool),
	}
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			fmt.Println("client connected.")
			h.server.clients = append(h.server.clients, new(Client))
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		case m := <-h.broadcast:
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					delete(h.connections, c)
					close(c.send)
				}
			}
		}
	}
}

// update gives each client connected to h its next state.
func (h *hub) update() {
	//send each channel its client state
	var state []byte
	for c := range h.connections {
		state = state[:0]
		view := c.client.GetMessage()
		for _, col := range view {
			for i := 0; i < len(col); i++ {
				state = append(state, col[i])
			}
		}
		c.ws.WriteMessage(websocket.BinaryMessage, []byte(state))
	}
}
