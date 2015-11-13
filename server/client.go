package main

import (
	"github.com/gummyworm/gogue/o/game"
)

// Client is the structure containing information about each player client.
type Client struct {
	X, Y   int //location of the client viewport's upper left corner
	W, H   int //width and height of viewport
	player game.Player
}

// NewClient creates a new client and returns a pointer to it.
func NewClient() *Client {
	var client Client
	client.H = 25
	client.W = 80
	client.player = game.Player{}
	client.player.Selection = append(client.player.Selection, game.Worm{})
	return &client
}

func (c *Client) GetMessage() (msg [80][25]byte) {
	return c.GetView()
}

// GetView returns the scene as seen by the camera of c.
func (c *Client) GetView() (ret [80][25]byte) {
	for i := 0; i < 25; i++ {
		for j := 0; j < 80; j++ {
			ret = c.player.See()
			ret[0][0] = 0x12
		}
	}
	return ret
}
