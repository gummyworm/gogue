package game

import (
	"fmt"
	"github.com/beefsack/go-astar"
)

type Tile struct {
	Cost int
	X, Y int
	Type int32
	room *Room
}

type Room struct {
	w, h  int      //width and height of the world map
	tiles [][]Tile //the tiles that compose the world map
}

// NewRoom creates a new room with dimensions wxh
func NewRoom(w, h int) *Room {
	r := &Room{w: w, h: h}
	r.tiles = make([][]Tile, h)
	for row := range r.tiles {
		r.tiles[row] = make([]Tile, w)
	}
	return r
}

// GetTile returns the tile at (x, y)
func (r *Room) GetTile(x, y int) Tile {
	fmt.Print(r.tiles[y][x])
	return r.tiles[y][x]
}

// SetTile sets the tile at (x, y) to the given tile type.
func (r *Room) SetTile(x, y int, Type int32) {
	r.tiles[y][x] = Tile{X: x, Y: y, Type: Type, room: r, Cost: 1}
}

// PathNeighbors returns the neighbors of the Tile
func (t Tile) PathNeighbors() []astar.Pather {
	x := []int{-1, 0, 1}
	y := []int{-1, 0, 1}
	var neighbors = []astar.Pather{}

	for i := range y {
		for j := range x {
			if t.X+x[j] >= 0 && t.X+x[j] < t.room.w && t.Y+y[i] >= 0 &&
				t.Y+y[i] < t.room.h {
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
