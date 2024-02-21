package entity

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	name  string
	x, y  float64
	Color *color.RGBA
	speed float64
}

type SpriteTexture struct {
}

func NewEntity(name string, x, y float64, c *color.RGBA) *Entity {
	return &Entity{
		name:  name,
		x:     x,
		y:     y,
		Color: c,
		speed: 16.0,
	}
}

func (e *Entity) SetSpeed(s float64) {
	e.speed = s
}

func (e *Entity) Update() error {
	return nil
}

func (e *Entity) Draw(screen *ebiten.Image) {
	// draw the entity on the screen
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			screen.Set(int(e.x)+i, int(e.y)+j, *e.Color)
		}
	}
}

func (e *Entity) MoveTowards(target *Entity) {
	dx := target.x - e.x
	dy := target.y - e.y
	dist := e.speed / (dx*dx + dy*dy)
	e.x += dx * dist
	e.y += dy * dist
}
