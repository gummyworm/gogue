package game

import (
	"github.com/beefsack/go-astar"
)

//Tile is a struct representing one unit of space in the game world.
type Tile struct {
	Glyph int32 //the glyph that this tile is rendered as.
	Type  int32
	X, Y  int
	Cost  int //the difficulty of the terrain (# of AP needed to move on it)
}

//World is a struct containing information about the game map.
type World struct {
	W, H  int
	tiles [][]Tile
}

//world is the active game world (map).
var world World

//GenWorld initializes the global world (map).
func GenWorld(w, h int) {
	world.tiles = make([][]Tile, w)
	world.H = h
	world.W = w
	for i := range world.tiles {
		world.tiles[i] = make([]Tile, h)
		for j := 0; j < h; j++ {
			world.tiles[i][j].X = i
			world.tiles[i][j].Y = j
		}
	}
}

//GetTile returns the tile at (x, y) in the game world or nil if there is none
func GetTile(x, y int) *Tile {
	if x >= 0 && x < world.W && y >= 0 && y < world.H {
		return &world.tiles[y][x]
	}
	return nil
}

//Render returns a wxh render of the world with the upper-left corner at (x, y).
func Render(x, y, w, h int, flags int) [][]byte {
	var ren = make([][]byte, w)
	for i := range ren {
		ren[i] = make([]byte, h)
	}
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			ren[i][j] = byte(world.tiles[i][j].Glyph)
		}
	}
	return ren
}

//SetTile sets the tile at (x, y) to the type described by id.
func SetTile(x, y int, id int32) {
	if x >= 0 && x < world.W && y >= 0 && y < world.H {
		//XXX base cost on type-ID
		world.tiles[y][x].Type = id
		world.tiles[y][x].Glyph = id
		world.tiles[y][x].Cost = 1
	}
}

func (t *Tile) PathNeighbors() []astar.Pather {
	var a []astar.Pather
	x := []int{-1, 0, 1}
	y := []int{-1, 0, 1}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			t := GetTile(t.X+x[i], t.Y+y[j])
			if t != nil {
				a = append(a, t)
			}
		}
	}
	return a
}

func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	tile := to.(*Tile)
	return float64(tile.Cost)
}

func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}
