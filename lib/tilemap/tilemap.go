package tilemap

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type TileMap struct {
	FilePath string
	Width    int
	Height   int
	Tiles    [][]Tile
}

type Tile struct {
	Width  int
	Height int
	Type   *TileType
}

type TileType struct {
	Name    string
	Texture *ebiten.Image
}

func NewTileMap(width, height int, mapPath string) *TileMap {
	tm := &TileMap{FilePath: mapPath}

	// Generate a new map and save it to the file
	tm.Width = width
	tm.Height = height
	tm.Tiles = make([][]Tile, tm.Width)
	for i := range tm.Tiles {
		tm.Tiles[i] = make([]Tile, tm.Height)
	}

	// Use Wave Function Collapse to set the type of each tile in the map
	// ...

	return tm
}
