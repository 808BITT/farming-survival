package tilemap

type TileMap struct {
	FilePath string
	Width    int
	Height   int
	Tiles    [][]Tile
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
