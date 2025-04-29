package point

import "math"

type Point struct {
	X float64
	Y float64
	Z float64
}

func (p Point) Rad() Point {
	return Point{
		X: p.X * math.Pi / 180,
		Y: p.Y * math.Pi / 180,
		Z: p.Z * math.Pi / 180,
	}
}
