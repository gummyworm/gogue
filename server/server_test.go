package main

import (
	"testing"
)

func TestServer(t *testing.T) {
	var ws WebServer
	ws.Init()
	ws.Start()
}
