package collision

import "fs/lib/phys"

// Hitbox is a hitbox
type Hitbox struct {
	Type     HitboxType
	Width    float64
	Height   float64
	Radius   float64
	Vertices []phys.Vec2
}

// HitboxType is the type of hitbox
type HitboxType int

const (
	// HitboxTypeCircle is a circle hitbox
	HitboxTypeCircle HitboxType = iota
	// HitboxTypeRectangle is a rectangle hitbox
	HitboxTypeRectangle
	// HitboxTypePolygon is a polygon hitbox
	HitboxTypePolygon
)

// NewHitbox creates a new hitbox
func NewRectHitbox(w, h float64) Hitbox {
	return Hitbox{
		Type:   HitboxTypeRectangle,
		Width:  w,
		Height: h,
	}
}

// NewCircleHitbox creates a new hitbox
func NewCircleHitbox(r float64) Hitbox {
	return Hitbox{
		Type:   HitboxTypeCircle,
		Radius: r,
	}
}

// NewPolygonHitbox creates a new hitbox
func NewPolygonHitbox(vertices []phys.Vec2) Hitbox {
	return Hitbox{
		Type:     HitboxTypePolygon,
		Vertices: vertices,
	}
}

// Collides checks if two hitboxes collide
func (h Hitbox) Collides(h2 Hitbox, p1, p2 phys.Vec2) bool {
	switch h.Type {
	case HitboxTypeCircle:
		switch h2.Type {
		case HitboxTypeCircle:
			return h.collidesCircleCircle(h2, p1, p2)
		case HitboxTypeRectangle:
			return h.collidesCircleRect(h2, p1, p2)
		case HitboxTypePolygon:
			return h.collidesCirclePolygon(h2, p1, p2)
		}
	case HitboxTypeRectangle:
		switch h2.Type {
		case HitboxTypeCircle:
			return h2.collidesCircleRect(h, p2, p1)
		case HitboxTypeRectangle:
			return h.collidesRectRect(h2, p1, p2)
		case HitboxTypePolygon:
			return h.collidesRectPolygon(h2, p1, p2)
		}
	case HitboxTypePolygon:
		switch h2.Type {
		case HitboxTypeCircle:
			return h2.collidesCirclePolygon(h, p2, p1)
		case HitboxTypeRectangle:
			return h2.collidesRectPolygon(h, p2, p1)
		case HitboxTypePolygon:
			return h.collidesPolygonPolygon(h2, p1, p2)
		}
	}
	return false
}

func (h Hitbox) collidesCircleCircle(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return p1.Distance(p2) < h.Radius+h2.Radius
}

func (h Hitbox) collidesCircleRect(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}

func (h Hitbox) collidesCirclePolygon(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}

func (h Hitbox) collidesRectRect(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}

func (h Hitbox) collidesRectPolygon(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}

func (h Hitbox) collidesPolygonPolygon(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}

func (h Hitbox) collidesPolygonRect(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}

func (h Hitbox) collidesPolygonCircle(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}

func (h Hitbox) collidesRectCircle(h2 Hitbox, p1, p2 phys.Vec2) bool {
	return false
}
