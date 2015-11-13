package game

import (
	"github.com/beefsack/go-astar"
	"testing"
)

func makeWorld(w, h int) *Room {
	room := NewRoom(w, h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			room.SetTile(j, i, 0x20)
		}
	}
	return room
}

func TestPathing(t *testing.T) {
	t.Log("creating new room")
	testRoom := makeWorld(100, 100)
	from := testRoom.GetTile(0, 0)
	to := testRoom.GetTile(15, 50)
	t.Log("pathing...")
	path, dist, found := astar.Path(from, to)
	if found {
		t.Logf("found path\ndistance: %v\npath:\n%v\n", dist, path)
	} else {
		t.Log("no path found")
	}
}
