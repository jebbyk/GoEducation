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

// inverses each component of a vector (components should not be grater than 1.0)
func (v *Vector2) Inverse(x bool, y bool) *Vector2 {
	if x {
		v.x = 1.0 - v.x
	}
	if y {
		v.y = 1.0 - v.y
	}
	return v
}

func (v *Vector2) GetX() float64 {
	return v.x
}

func (v *Vector2) GetY() float64 {
	return v.y
}
