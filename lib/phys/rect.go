package phys

type Rect struct {
	X      float64
	Y      float64
	Width  float64
	Height float64
}

func NewRect(x, y, w, h float64) *Rect {
	return &Rect{
		X:      x,
		Y:      y,
		Width:  w,
		Height: h,
	}
}

func (r Rect) Contains(p Vec2) bool {
	return p.X >= r.X && p.X <= r.X+r.Width && p.Y >= r.Y && p.Y <= r.Y+r.Height
}

func (r Rect) Intersects(r2 Rect) bool {
	return r.X < r2.X+r2.Width && r.X+r.Width > r2.X && r.Y < r2.Y+r2.Height && r.Y+r.Height > r2.Y
}

func (r Rect) Center() Vec2 {
	return NewVec2(r.X+r.Width/2, r.Y+r.Height/2)
}

func (r Rect) TopLeft() Vec2 {
	return NewVec2(r.X, r.Y)
}

func (r Rect) TopRight() Vec2 {
	return NewVec2(r.X+r.Width, r.Y)
}

func (r Rect) BottomLeft() Vec2 {
	return NewVec2(r.X, r.Y+r.Height)
}

func (r Rect) BottomRight() Vec2 {
	return NewVec2(r.X+r.Width, r.Y+r.Height)
}

func (r Rect) Vertices() []Vec2 {
	return []Vec2{
		r.TopLeft(),
		r.TopRight(),
		r.BottomRight(),
		r.BottomLeft(),
	}
}
