package graphs

import . "mathMethods/geometry/types"

type Graph struct {
	position *Vector2
	scale    *Vector2
}

func (g *Graph) Init(position *Vector2, scale *Vector2) *Graph {
	g.position = position
	g.scale = scale

	return g
}

func (g *Graph) VisualizeFloats(values []float64) {

}
