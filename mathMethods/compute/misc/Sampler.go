package misc

func Downsample(values []float64, samplesAmount int) []float64 {

	samplesPerStep := len(values) / samplesAmount

	samples := []float64{}

	for i := 0; i < samplesAmount; i++ {
		sum := 0.0
		count := 0
		for j := 0; j < samplesPerStep; j++ {
			elementNumber := i*samplesPerStep + j
			if elementNumber < len(values) {
				sum += values[elementNumber] //sum all the values form current array part
				count++
			}
			continue // prevent counting over a length of values array length
		}
		sum /= float64(samplesPerStep) // find mid values
		samples = append(samples, sum)
	}
	return samples
}
