package main

import (
	"fmt"
	"mathMethods/compute/random"
	. "mathMethods/geometry/types"
	. "mathMethods/interact/prompt"
	. "mathMethods/visualize/buffers/console"
	. "mathMethods/visualize/graphs"
)

func main() {
	generator := new(random.Generator).Init(0)
	prompter := new(Prompter).Init("\n", "Incorrect input!")

	iterations := prompter.RequestInteger("Enter iterations amount: ")

	var values []float64

	for i := 0; i < iterations; i++ {
		values = append(values, generator.Fgen())
	}

	buffer := new(Buffer).Init(32, 128)

	graphPos := new(Vector2).Set(0.1, 0.9)
	graphScale := new(Vector2).Set(0.9, 0.9)
	graph := new(Graph).Init(graphPos, graphScale)

	fmt.Println("Success!")
}
