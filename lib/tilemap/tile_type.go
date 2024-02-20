package tilemap

import "github.com/hajimehoshi/ebiten/v2"

type TileTexture struct {
	Texture *ebiten.Image
	Edge    TileEdges
}

func NewTileTexture(name string, texture *ebiten.Image, edges TileEdges) TileTexture {
	return TileTexture{
		Texture: texture,
		Edge:    edges,
	}
}
