package main

import (
	"fmt"
)

// Server is the structure used to contain the server-side game state.
type Server struct {
	tickCnt int
}

func MakeScene() {
}

// Init prepares the server for running.
func (server *Server) Init() {
	fmt.Println("(((O)))")
	server.tickCnt = 0 //initialize server tick count
	MakeScene()        //setup the initial scene
	fmt.Println("server initialized")
}

// Tick runs one "tick" (update) of the engine.
func (server *Server) Tick() {
	server.tickCnt++
}
