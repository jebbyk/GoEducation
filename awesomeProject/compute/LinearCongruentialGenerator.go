package compute

const MAX = 2 ^ 32 // TODO check arch somehow

type Generator struct {
	next       int
	multiplier int
	delta      int
}

func (g Generator) init(seed int) Generator {
	g.next = seed
	g.multiplier = 1103515245
	g.delta = 12345

	return g
}

func (g Generator) gen() int {
	g.next = (g.next*g.multiplier + g.delta) % (MAX + 1)
	return g.next
}

func (g Generator) fgen() float32 {
	return float32(g.gen()) / MAX
}
