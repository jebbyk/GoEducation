package main

import (
	"fmt"
	"math"
	"mathMethods/compute/misc"
	"mathMethods/compute/random"
	. "mathMethods/geometry/types"
	. "mathMethods/interact/prompt"
	. "mathMethods/visualize/buffers/console"
	. "mathMethods/visualize/graphs"
)

func countProbabilities(values []float64, stepsAmount int) []float64 {
	min := misc.FindMinFloat(values)
	max := misc.FindMaxFloat(values)

	rng := max - min

	stepLength := rng / float64(stepsAmount)

	probabilities := []float64{}

	for i := 0; i < stepsAmount; i++ {
		localStepRangeMin := float64(i) * stepLength
		localStepRangeMax := localStepRangeMin + stepLength
		count := 0

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

func generateEvenDistribution(iterations int, generator *random.Generator) []float64 {
	var values []float64
	for i := 0; i < iterations; i++ {
		values = append(values, generator.Fgen())
	}

	return values
}
func calcEstimatedEvenProbabilities(min float64, max float64, stepsAmount int) []float64 {
	rng := max - min
	stepLength := rng / float64(stepsAmount)

	probabilities := []float64{}

	for i := 0; i < stepsAmount; i++ {
		lft := min + float64(i)*stepLength
		rgt := lft + stepLength

		probabilities = append(probabilities, rgt-lft)
	}

	return probabilities
}

func generateExpDistribution(exponent float64, iterations int, generator *random.Generator) []float64 {
	var values []float64
	for i := 0; i < iterations; i++ {
		value := -exponent * math.Log10(generator.Fgen())
		values = append(values, value)
	}

	return values
}
func calcEstimatedExpProbabilities(min float64, max float64, stepsAmount int) []float64 {
	rng := max - min
	stepLength := rng / float64(stepsAmount)

	probabilities := []float64{}

	for i := 0; i < stepsAmount; i++ {
		lft := min + float64(i)*stepLength
		rgt := lft + stepLength

		prob := math.Exp(-lft) - math.Exp(-rgt)
		probabilities = append(probabilities, prob)
	}
	return probabilities
}

func generateNormDistribution(Q float64, iterations int, generator *random.Generator) []float64 {
	values := []float64{}
	for i := 0; i < iterations/2; i++ {

		u0 := generator.Fgen()
		u1 := generator.Fgen()

		v0 := u0*2.0 - 1.0
		v1 := u1*2.0 - 1.0

		S := v0*v0 + v1*v1

		if S >= 1.0 {
			i--
			continue
		}

		x0 := v0 * math.Sqrt((-2.0*math.Log10(S))/S)
		x1 := v1 * math.Sqrt((-2.0*math.Log10(S))/S)

		x0 = x0 * Q
		x1 = x1 * Q

		values = append(values, x0, x1)
	}
	return values
}
func idkHowToNameThisShit(value float64) float64 {
	exp := math.Exp(-0.5 * (value * value)) // https://ru.wikipedia.org/wiki/%D0%9D%D0%BE%D1%80%D0%BC%D0%B0%D0%BB%D1%8C%D0%BD%D0%BE%D0%B5_%D1%80%D0%B0%D1%81%D0%BF%D1%80%D0%B5%D0%B4%D0%B5%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5
	fract := 1.0 / math.Sqrt(2.0*math.Pi)
	return fract * exp
}
func calcEstimatedNormProbabilities(min float64, max float64, stepsAmount int) []float64 {
	rng := max - min
	stepLength := rng / float64(stepsAmount)

	probabilities := []float64{}

	for i := 0; i < stepsAmount; i++ {
		lft := min + float64(i)*stepLength
		rgt := lft + stepLength

		prob := idkHowToNameThisShit(((lft - rgt) / rgt) * (lft - rgt))

		probabilities = append(probabilities, prob)
	}
	return probabilities
}

func pearsonCriterion(estimated []float64, actual []float64) float64 {
	estimatedLength := len(estimated)
	actualLength := len(actual)

	if estimatedLength != actualLength {
		return 999999999999.9 // idk how to "throw exception" in go :P
	}

	result := 0.0
	for i := 0; i < estimatedLength; i++ {
		result += math.Pow((actual[i]/float64(actualLength))-estimated[i], 2.0) / estimated[i]
	}

	return result
}

func displayValues(values []float64, buffer *Buffer, posX float64, posY float64, scaleX float64, scaleY float64, steps int) {

	probabilities := countProbabilities(values, steps)
	graphPos := new(Vector2).Set(posX, posY)
	graphScale := new(Vector2).Set(scaleX, scaleY)
	graph := new(Graph).Init(graphPos, graphScale, STYLE_BARS)
	graph.Draw(probabilities, buffer)

}

func main() {
	generator := new(random.Generator).Init(0)
	prompter := new(Prompter).Init("\n", "Incorrect input!")
	buffer := new(Buffer).Init(110, 350)

	iterations := prompter.RequestInteger("Enter iterations amount: ")
	//iterations := 10
	exponent := prompter.RequestFloat("Enter exponent grater than 0: ")
	Q := prompter.RequestFloat("Enter Q between grater than 0: ")

	norm := []float64{}

	//for iterations := 2; iterations < 1000000; iterations += 2 { // add 2 in each iteration cuz normDistribution generates 2 numbers at once
	even := generateEvenDistribution(iterations, generator)
	estimatedEvenProbabilities := calcEstimatedEvenProbabilities(misc.FindMinFloat(even), misc.FindMaxFloat(even), 20)
	actualEvenProbabilities := countProbabilities(even, 20)
	evenPearsonCriterion := pearsonCriterion(estimatedEvenProbabilities, actualEvenProbabilities)
	displayValues(even, buffer, 0.0, 0.0, 0.3, 8.0, 100)

	exp := generateExpDistribution(exponent, iterations, generator)
	estimatedExpProbabilities := calcEstimatedExpProbabilities(misc.FindMinFloat(exp), misc.FindMaxFloat(exp), 20)
	actualExpProbabilities := countProbabilities(exp, 20)
	expPearsonCriterion := pearsonCriterion(estimatedExpProbabilities, actualExpProbabilities)
	displayValues(exp, buffer, 0.35, 0.0, 0.3, 8.0, 100)

	norm = generateNormDistribution(Q, iterations, generator)
	for i := 0; i < iterations; i++ {
		norm[i] += 2.0 //Offset them so they can be displayed on my shitty graph :)
	}
	estimatedNormProbabilities := calcEstimatedNormProbabilities(misc.FindMinFloat(norm), misc.FindMaxFloat(norm), 20)
	actualNormProbabilities := countProbabilities(norm, 20)
	normPearsonCriterion := pearsonCriterion(estimatedNormProbabilities, actualNormProbabilities)
	displayValues(norm, buffer, 0.7, 0.0, 0.3, 8.0, 100)

	fmt.Println("=========================================================================================================================================================================================================")
	buffer.Print()
	fmt.Println("=========================================================================================================================================================================================================")

	fmt.Println("even: ", evenPearsonCriterion)
	fmt.Println("exp: ", expPearsonCriterion)
	fmt.Println("norm: ", normPearsonCriterion)
	buffer.Clear()

	//time.Sleep(100 * time.Millisecond)
	//}
}
