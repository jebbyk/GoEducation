package main

import (
	"fmt"
	"mathMethods/compute/misc"
	"mathMethods/compute/random"
	. "mathMethods/geometry/types"
	. "mathMethods/interact/prompt"
	. "mathMethods/visualize/buffers/console"
	. "mathMethods/visualize/graphs"
)

func countProbabilities(values []float64, steps int) []float64 {
	min := misc.FindMinFloat(values)
	max := misc.FindMaxFloat(values)

	//fmt.Println(min, max)

	rng := max - min //TODO take negative number into account later

	//fmt.Println(rng)

	stepLength := rng / float64(steps)

	//fmt.Println(stepLength)

	probabilities := []float64{}

	for i := 0; i < steps; i++ {
		localStepRangeMin := float64(i) * stepLength
		localStepRangeMax := localStepRangeMin + stepLength
		count := 0

		//fmt.Println(localStepRangeMin, localStepRangeMax)

		for j := 0; j < len(values); j++ {
			if values[j] > localStepRangeMin && values[j] < localStepRangeMax {
				count++
			}
		}

		probability := float64(count) / float64(len(values))

		probabilities = append(probabilities, probability)

	}

	return probabilities
}

func main() {
	generator := new(random.Generator).Init(0)
	prompter := new(Prompter).Init("\n", "Incorrect input!")

	iterations := prompter.RequestInteger("Enter iterations amount: ")

	var values []float64

	for i := 0; i < iterations; i++ {
		values = append(values, generator.Fgen())
		//fmt.Println(generator.Fgen())
	}

	probabilities := countProbabilities(values, 20)

	//fmt.Println(probabilities)

	buffer := new(Buffer).Init(32, 100)

	graphPos := new(Vector2).Set(0.05, 0.05)
	graphScale := new(Vector2).Set(0.9, 0.9)
	graph := new(Graph).Init(graphPos, graphScale)

	graph.DisplayF64(probabilities, buffer)

	fmt.Println("=================================================================================================================")
	buffer.Print()
	fmt.Println("=================================================================================================================")

	fmt.Println("Success!")
}
