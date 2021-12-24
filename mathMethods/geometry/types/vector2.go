package types

import "math"

type Vector2 struct {
	x float64
	y float64
}

// set exact values of a vector
func (v *Vector2) Set(x float64, y float64) *Vector2 {
	v.x = x
	v.y = y
	return v
}

// add two vectors
func (v0 *Vector2) Add(v1 *Vector2) *Vector2 {
	v0.x += v1.x
	v0.y += v1.y

	return v0
}

// multiply two vectors
func (v0 *Vector2) Mul(v1 *Vector2) *Vector2 {
	v0.x *= v1.x
	v0.y *= v1.y

	return v0
}

// find magnitude of two vectors
func (v0 *Vector2) SqrMagnitude() float64 {
	return v0.x*v0.x + v0.y*v0.y
}

// calculate precise length of two vectors
func (v0 *Vector2) Len() float64 {
	return math.Sqrt(v0.SqrMagnitude())
}
