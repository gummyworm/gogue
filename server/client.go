package main

import (
	"github.com/gummyworm/gogue/game"
)

// Client is the structure containing information about each player client.
type Client struct {
	view   [80 * 25]byte
	player *game.Player
}

const UPDATE_MSG_LEN = 80 * 25

// NewClient creates a new client and returns a pointer to it.
func NewClient() *Client {
	var client Client
	return &client
}

// GetView returns the scene as seen by the camera of c.
func (c *Client) GetMessage() []byte {
	msg := make([]byte, UPDATE_MSG_LEN)

	// get the "view" for this player
	view := c.player.See()
	return msg
	for i := 0; i < 25; i++ {
		for j := 0; j < 80; j++ {
			msg[i*80+j] = byte(view[i][j])
		}
	}
	return msg
}
