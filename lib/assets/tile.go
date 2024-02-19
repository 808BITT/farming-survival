package assets

import (
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Tile struct {
	Id    uuid.UUID
	Name  string
	Img   *ebiten.Image
	North Edge
	West  Edge
	South Edge
	East  Edge
}

func NewTile(name string, image *ebiten.Image, north, west, south, east Edge) Tile {
	return Tile{
		Id:    uuid.New(),
		Name:  name,
		Img:   image,
		North: north,
		West:  west,
		South: south,
		East:  east,
	}
}

func (t Tile) Equals(other Tile) bool {
	if t.Id != other.Id {
		return false
	}

	if t.Name != other.Name {
		return false
	}

	if t.North != other.North {
		return false
	}

	if t.West != other.West {
		return false
	}

	if t.South != other.South {
		return false
	}

	if t.East != other.East {
		return false
	}

	return true
}

func (t Tile) String() string {
	return t.Name
}

type Edge string

func NewEdge(name string) Edge {
	return Edge(name)
}
