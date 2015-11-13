package game

import (
	"github.com/beefsack/go-astar"
	"testing"
)

func makeTestWorld(w, h int) {
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			SetTile(i, j, 0x40)
		}
	}
}

func TestTile(t *testing.T) {
	const x = 3
	const y = 4
	const id = 42

	GenWorld(100, 100)

	SetTile(x, y, id)
	tile := GetTile(x, y)

	//XXX: verify match
	if tile.Glyph != id {
		t.Error("SetTile failed")
	}
}

func TestWorld(t *testing.T) {
	GenWorld(100, 100)
	makeTestWorld(100, 100)

	from := GetTile(0, 0)
	to := GetTile(9, 10)

	_, dist, found := astar.Path(to, from)
	if !found {
		t.Log("no path found")
	} else {
		t.Log("path found. total distance: ", dist)
	}
}

func BenchmarkWorld(b *testing.B) {
	b.Log("testing world")
	makeTestWorld(b.N, b.N)

	from := GetTile(0, 0)
	to := GetTile(b.N-1, b.N-1)
	for i := 0; i < b.N; i++ {
		astar.Path(to, from)
	}
}
