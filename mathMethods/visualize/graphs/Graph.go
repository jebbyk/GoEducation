package graphs

import (
	. "mathMethods/geometry/types"
	"mathMethods/visualize/buffers/console"
)

const STYLE_DOTS = 0
const STYLE_BARS = 1
const STYLE_LINES = 2

type Graph struct {
	position *Vector2
	scale    *Vector2

	style int // dots or bars so far
}

func (g *Graph) Init(position *Vector2, scale *Vector2, style int) *Graph {
	g.position = position
	g.scale = scale

	return g
}

func (g *Graph) Draw(values []float64, buffer *console.Buffer) {
	if g.style == STYLE_DOTS {
		g.drawDots(values, buffer)
	} else if g.style == STYLE_BARS {
		g.drawBars(values, buffer)
	} else if g.style == STYLE_LINES {
		//TODO
	}
}

func (g *Graph) drawDots(values []float64, buffer *console.Buffer) {
	valuesAmount := float64(len(values))
	for i := 0; i < len(values); i++ {
		xPos0 := float64(i) / valuesAmount
		yPos0 := values[i] //TODO accept numbers outside of range [0...1] too, autoscale graph

		barTopLeft := new(Vector2).Set(xPos0, yPos0).Mul(g.scale).Add(g.position).Inverse(false, true)

		xPos1 := float64(i+1) / valuesAmount
		yPos1 := 0.0
		barBottomRight := new(Vector2).Set(xPos1, yPos1).Mul(g.scale).Add(g.position).Inverse(false, true)

		buffer.DrawRect(barTopLeft, barBottomRight, '+')
	}
}

func (g *Graph) drawBars(values []float64, buffer *console.Buffer) {
	valuesAmount := float64(len(values))
	for i := 0; i < len(values); i++ {
		xPos := float64(i) / valuesAmount
		yPos := values[i] //TODO accept numbers outside of range [0...1] too, autoscale graph

		barTopLeft := new(Vector2).Set(xPos, yPos).Mul(g.scale).Add(g.position).Inverse(false, true)

		buffer.SetPixelRelative(barTopLeft, '+')
	}
}
