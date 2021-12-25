package main

import (
	"fmt"
	"mathMethods/compute/random"
	"mathMethods/interact/prompt"
	"mathMethods/visualize/graphs"
)

func main() {
	generator := new(random.Generator).Init(0)
	prompter := new(prompt.Prompter).Init("\n", "Incorrect input!")

	iterations := prompter.RequestInteger("Enter iterations amount: ")

	for i := 0; i < iterations; i++ {
		fmt.Println(generator.Fgen())
	}

	graph := new(graphs.Graph).Init()

	fmt.Println("Success!")
}
