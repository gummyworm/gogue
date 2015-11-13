package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/gummyworm/gogue/game"
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

//XXX delete
var testRoom *game.Room
var player *game.Player

func newHub(server *WebServer) *hub {
	//XXX delete
	game.Init()
	testRoom = game.NewRoom("Room", 100, 100)
	game.State.AddRoom(testRoom)

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
			//XXX look up player using client-provided information
			player = &game.Player{Name: "p1", X: 0, Y: 0, W: 80, H: 25}
			c.client = &Client{player: player}
			game.State.AddPlayer(player)
			game.State.SetPlayer("p1", "Room")
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
	for c := range h.connections {
		msg := c.client.GetMessage()
		c.ws.WriteMessage(websocket.BinaryMessage, msg)
	}
}
