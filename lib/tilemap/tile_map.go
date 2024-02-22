package tilemap

type TileMap struct {
	Width  int
	Height int
	Tiles  [][]Tile
}

func NewTileMap(width, height int) *TileMap {
	tm := &TileMap{
		Width:  width,
		Height: height,
		Tiles:  make([][]Tile, width),
	}

	for i := range tm.Tiles {
		tm.Tiles[i] = make([]Tile, tm.Height)
	}

	return tm
}
