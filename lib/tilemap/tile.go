package tilemap

import (
	"github.com/google/uuid"
)

type Tile struct {
	Id          uuid.UUID
	Name        string
	Texture     TileTexture
	Width       int
	Height      int
	Probability float64
}

func NewTile(tt TileTexture, p float64) Tile {
	return Tile{
		Id:          uuid.New(),
		Texture:     tt,
		Width:       tt.Img.Bounds().Dx(),
		Height:      tt.Img.Bounds().Dy(),
		Probability: p,
	}
}

func (t Tile) Equals(other Tile) bool {
	return t.Id == other.Id
}

func (t Tile) String() string {
	return t.Name
}
