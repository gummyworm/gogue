package main

import (
	"flag"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	addr     = flag.String("localhost", ":8080", "O")
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
	tickSpeed = 500 * time.Millisecond // time between each server tick
)

type wsHandler struct {
	h *hub
}

// ServeHTTP is the websocket handler.
func (wsh wsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{send: make(chan []byte, 256), ws: ws, h: wsh.h, client: new(Client)}
	c.h.register <- c
	defer func() { c.h.unregister <- c }()
	go c.writer()
	c.reader()
}

// WebServer is a server that talks to clients over WebSockets.
type WebServer struct {
	Server
}

// Init initializes the server and creates a network connection.
func (ws *WebServer) Init() {
	ws.Server.Init()
}

// Start begins execution of the webserver.
// This function does not return.
func (ws *WebServer) Start() {
	//create the hub- hanlde connections & disconnections
	hub := newHub(ws)
	go hub.run()

	//serve WebSocket and static HTTP
	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("static"))))
	http.Handle("/ws", wsHandler{h: hub})
	go http.ListenAndServe(":8080", nil)

	//run main server loop- update the engine & send each client its new state.
	for {
		loopStartTime := time.Now()
		ws.Server.Tick()
		//if we're early to the next server tick, wait til its time to update.
		if time.Since(loopStartTime) > 0 {
			time.Sleep(tickSpeed)
		}
		hub.update() //send each client its state for this frame
	}
}
