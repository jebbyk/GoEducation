package random

const MAX = 151875
const MULT = 31
const DELTA = 1283

// contains all the data needed to generate pseudorandom numbers
type Generator struct {
	next       int
	multiplier int
	delta      int
}

// initialize a Generator
func (g *Generator) Init(seed int) *Generator {
	g.next = seed
	g.multiplier = MULT
	g.delta = DELTA

	return g
}

// generates a pseudorandom integer
func (g *Generator) Gen() int {
	g.next = (g.next*g.multiplier + g.delta) % MAX
	return g.next
}

// generates a pseudorandom float32
func (g *Generator) Fgen() float64 {
	return float64(g.Gen()) / MAX
}
