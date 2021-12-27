package misc

func FindMinFloat(values []float64) float64 {
	min := 999999999999999.9 //TODO something with it later but now im lazy :P
	for i := 0; i < len(values); i++ {
		if values[i] < min {
			min = values[i]
		}
	}
	return min
}

func FindMaxFloat(values []float64) float64 {
	max := -99999999999999.9
	for i := 0; i < len(values); i++ {
		if values[i] > max {
			max = values[i]
		}
	}
	return max
}
