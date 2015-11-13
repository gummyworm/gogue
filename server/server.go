package main

import (
	"fmt"
	"github.com/gummyworm/gogue/tv"
)

// Server is the structure used to contain the server-side game state.
type Server struct {
	tickCnt int
	clients []*Client //slice of each connected client
}

// getClientMessages returns the messages prepared for each client.
func (s *Server) getClientMessages() {
	//build each client message
	for _, c := range s.clients {
		c.GetMessage()
	}
}

func MakeScene() {
}

// Init prepares the server for running.
func (server *Server) Init() {
	fmt.Println("(((O)))")

	server.tickCnt = 0 //initialize server tick count
	tv.Init()          //initialize subsystems
	MakeScene()        //setup the initial scene
	tv.Start()         //start all the entities

	fmt.Println("server initialized")
}

// Tick runs one "tick" (update) of the engine.
func (server *Server) Tick() {
	tv.Update()
	server.tickCnt++
}
