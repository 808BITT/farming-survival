package wfc

import (
	"fs/lib/assets"
)

type CollapseTile struct {
	Tile          *assets.Tile
	PossibleTiles []*assets.Tile
}

func NewCollapseTile(tile *assets.Tile) *CollapseTile {
	return &CollapseTile{
		Tile:          tile,
		PossibleTiles: nil,
	}
}
