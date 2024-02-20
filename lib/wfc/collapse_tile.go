package wfc

import "fs/lib/tilemap"

type CollapseTile struct {
	Tile          *tilemap.Tile
	PossibleTiles []*tilemap.Tile
}

func NewCollapseTile(tile *tilemap.Tile) *CollapseTile {
	return &CollapseTile{
		Tile:          tile,
		PossibleTiles: nil,
	}
}
