package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	x, y  float64
	eType *EntityType
	speed float64
}

type EntityType struct {
	Name  string
	Color *color.Color
}

func NewEntity(x, y float64, eType *EntityType) *Entity {
	return &Entity{
		x:     x,
		y:     y,
		eType: eType,
		speed: 1,
	}
}

func (e *Entity) Draw(screen *ebiten.Image) {
	// draw the entity on the screen
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			screen.Set(int(e.x)+i, int(e.y)+j, *e.eType.Color)
		}
	}
}

func (e *Entity) Update(other []*Entity) {
	for _, entity := range other {
		if entity == e {
			continue
		}
		if entity.eType.Name == e.eType.Name {
			e.moveTowards(entity)
		}
	}
}

func (e *Entity) moveTowards(target *Entity) {
	dx := target.x - e.x
	dy := target.y - e.y
	dist := e.speed / (dx*dx + dy*dy)
	e.x += dx * dist
	e.y += dy * dist
}
