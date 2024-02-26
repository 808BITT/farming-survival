package phys

import "math"

type Vec2 struct {
	X float64
	Y float64
}

func NewVec2(x, y float64) Vec2 {
	return Vec2{X: x, Y: y}
}

func (v Vec2) Normalize() Vec2 {
	length := math.Sqrt(v.X*v.X + v.Y*v.Y)
	return Vec2{X: v.X / length, Y: v.Y / length}
}

func (v Vec2) Magnitude() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vec2) Add(v2 Vec2) Vec2 {
	return Vec2{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vec2) Sub(v2 Vec2) Vec2 {
	return Vec2{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vec2) Mul(s float64) Vec2 {
	return Vec2{X: v.X * s, Y: v.Y * s}
}

func (v Vec2) Div(s float64) Vec2 {
	return Vec2{X: v.X / s, Y: v.Y / s}
}

func (v Vec2) Dot(v2 Vec2) float64 {
	return v.X*v2.X + v.Y*v2.Y
}

func (v Vec2) Cross(v2 Vec2) float64 {
	return v.X*v2.Y - v.Y*v2.X
}

func (v Vec2) Angle(v2 Vec2) float64 {
	return math.Acos(v.Dot(v2) / (v.Magnitude() * v2.Magnitude()))
}

func (v Vec2) Distance(v2 Vec2) float64 {
	return v.Sub(v2).Magnitude()
}

func (v Vec2) Project(v2 Vec2) Vec2 {
	return v.Mul(v.Dot(v2) / v.Dot(v))
}

func (v Vec2) Reflect(v2 Vec2) Vec2 {
	return v.Sub(v2.Mul(2 * v.Dot(v2) / v2.Dot(v2)))
}

func (v Vec2) AngleBetween(v2 Vec2) float64 {
	return math.Atan2(v2.Y-v.Y, v2.X-v.X) - math.Atan2(0, 1)
}

func (v Vec2) RotateAround(v2 Vec2, angle float64) Vec2 {
	return v.Sub(v2).Rotate(angle).Add(v2)
}

func (v Vec2) Scale(s float64) Vec2 {
	return Vec2{X: v.X * s, Y: v.Y * s}
}

func (v Vec2) Rotate(angle float64) Vec2 {
	cos := math.Cos(angle)
	sin := math.Sin(angle)
	return Vec2{X: v.X*cos - v.Y*sin, Y: v.X*sin + v.Y*cos}
}

func (v Vec2) Lerp(v2 Vec2, t float64) Vec2 {
	return v.Add(v2.Sub(v).Mul(t))
}
