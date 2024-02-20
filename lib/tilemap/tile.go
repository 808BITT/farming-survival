package tilemap

import (
	"github.com/google/uuid"
)

type Tile struct {
	Id          uuid.UUID
	Name        string
	Type        TileTexture
	Width       int
	Height      int
	Probability float64
}

func NewTile(tileType TileTexture, probability float64) Tile {
	return Tile{
		Id:          uuid.New(),
		Type:        tileType,
		Width:       tileType.Texture.Bounds().Dx(),
		Height:      tileType.Texture.Bounds().Dy(),
		Probability: probability,
	}
}

func (t Tile) Equals(other Tile) bool {
	return t.Id == other.Id
}

func (t Tile) String() string {
	return t.Name
}
