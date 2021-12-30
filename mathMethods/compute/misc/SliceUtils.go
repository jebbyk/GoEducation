package misc

func FindMinFloat(values []float64) int {
	min := 999999999999999.9 //TODO something with it later but now im lazy :P
	id := 0
	for i := 0; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
			id = i
		}
	}
	return id
}

func FindMaxFloat(values []float64) int {
	max := -99999999999999.9
	id := 0
	for i := 0; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
			id = i
		}
	}
	return id
}
