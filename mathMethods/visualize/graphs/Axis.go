package graphs

type Axis struct {
	name   string
	values []string
}

func (a *Axis) Init(name string, values []string) *Axis {
	a.name = name
	a.values = values

	return a
}
