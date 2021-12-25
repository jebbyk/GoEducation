package graphs

import "mathMethods/geometry/types"

type Graph struct {
	_type string //bars, lines
	xAxis *Axis
	yAxis *Axis
	scale *types.Vector2
}
