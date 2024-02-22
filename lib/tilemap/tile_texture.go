package tilemap

import "github.com/hajimehoshi/ebiten/v2"

type TileTexture struct {
	Img  *ebiten.Image
	Edge TileEdges
}

func NewTileTexture(texture *ebiten.Image, edges TileEdges) TileTexture {
	return TileTexture{
		Img:  texture,
		Edge: edges,
	}
}
