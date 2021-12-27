package graphs

import (
	. "mathMethods/geometry/types"
	. "mathMethods/visualize/buffers/console"
)

type Graph struct {
	position *Vector2
	scale    *Vector2
}

func (g *Graph) Init(position *Vector2, scale *Vector2) *Graph {
	g.position = position
	g.scale = scale

	return g
}

func (g *Graph) DisplayF64(values []float64, buffer *Buffer) {
	valuesAmount := float64(len(values))
	for i := 0; i < len(values); i++ {
		xPos := float64(i) / valuesAmount
		yPos := values[i] //TODO accept numbers outside of range [0...1] too, autoscale graph

		pos := new(Vector2).Set(xPos, yPos).Mul(g.scale).Add(g.position).Inverse(false, true)

		buffer.SetPixelRelative(pos, '+')
	}
}
