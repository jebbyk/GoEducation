package primitives

import "mathMethods/geometry/types"

type Point struct {
	coordinates types.Vector2
}

// sets new point coordinates
func (p *Point) SetCoordinates(x float64, y float64) *Point {
	p.coordinates.Set(x, y)
	return p
}

// adds a vector to current point coordinates
func (p *Point) Move(v *types.Vector2) *Point {
	p.coordinates.Add(v)
	return p
}
