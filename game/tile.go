package game

import (
	"github.com/beefsack/go-astar"
)

type Tile struct {
	Cost  int
	X, Y  int
	Type  int32
	Glyph int32
	room  *Room
}

const (
	TILE_NONE = iota
	TILE_PLAIN
	TILE_STONE
)

type Room struct {
	Name  string   //the room's name
	W, H  int      //width and height of the world map
	tiles [][]Tile //the tiles that compose the world map
}

// NewRoom creates a new room with dimensions wxh
func NewRoom(name string, w, h int) *Room {
	r := &Room{Name: name, W: w, H: h}
	r.tiles = make([][]Tile, h)
	for i := range r.tiles {
		r.tiles[i] = make([]Tile, w)
	}
	return r
}

// GetTile returns the tile at (x, y)
func (r *Room) GetTile(x, y int) Tile {
	if x >= 0 && x < r.W && y >= 0 && y <= r.H {
		return r.tiles[y][x]
	}
	return Tile{X: x, Y: y, Type: TILE_NONE, Glyph: ' ', room: r}
}

// SetTile sets the tile at (x, y) to the given tile type.
func (r *Room) SetTile(x, y int, Type int32) {
	r.tiles[y][x] = Tile{X: x, Y: y, Type: Type, Glyph: Type, room: r, Cost: 1}
}

// PathNeighbors returns the neighbors of the Tile
func (t Tile) PathNeighbors() []astar.Pather {
	x := []int{-1, 0, 1}
	y := []int{-1, 0, 1}
	var neighbors = []astar.Pather{}

	for i := range y {
		for j := range x {
			if t.X+x[j] >= 0 && t.X+x[j] < t.room.W && t.Y+y[i] >= 0 &&
				t.Y+y[i] < t.room.H {
				neighbors = append(neighbors,
					t.room.tiles[t.Y+y[i]][t.X+x[j]])
			}
		}
	}
	return neighbors
}

// PathNeighborCost returns the cost of moving to the given tile
func (t Tile) PathNeighborCost(to astar.Pather) float64 {
	return float64(t.Cost)
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	r := float64(absX + absY)
	return r
}
