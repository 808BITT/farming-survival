package tilemap

type TileEdges map[string]string

func NewTileEdges(north, east, south, west string) TileEdges {
	return TileEdges{
		"north": north,
		"west":  west,
		"south": south,
		"east":  east,
	}
}

func (te TileEdges) North() string {
	return te["north"]
}

func (te TileEdges) East() string {
	return te["east"]
}

func (te TileEdges) South() string {
	return te["south"]
}

func (te TileEdges) West() string {
	return te["west"]
}
