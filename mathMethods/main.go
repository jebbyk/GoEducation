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

func countProbabilities(values []float64, stepsAmount int, verbose bool) []float64 {
	min := values[misc.FindMinFloat(values)]
	max := values[misc.FindMaxFloat(values)]

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

	if verbose {
		maxProbId := misc.FindMaxFloat(probabilities)
		l := float64(maxProbId) * stepLength
		r := l + stepLength
		fmt.Println("MinNumber: ", min, ". MostFrequentPart: ", l, "-", r, ". MaxNumber: ", max, ". MostFreqPartProb: ", probabilities[maxProbId])
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

func generateNormDistribution(Q float64, A float64, iterations int, generator *random.Generator) []float64 {
	values := []float64{}
	for i := 0; i < iterations/2; i++ {

		u0 := generator.Fgen()
		u1 := generator.Fgen()

		v0 := (u0 * 2.0) - 1.0
		v1 := (u1 * 2.0) - 1.0

		S := v0*v0 + v1*v1

		if S >= 1.0 {
			i--
			continue
		}

		x0 := v0 * math.Sqrt((-2.0*math.Log10(S))/S)
		x1 := v1 * math.Sqrt((-2.0*math.Log10(S))/S)

		x0 = x0*Q + A
		x1 = x1*Q + A

		values = append(values, x0, x1)
	}
	return values
}
func normDistributionProbDensity(Q float64, A float64, value float64) float64 {
	expFract0 := (value - A) * (value - A)
	expFract1 := 2 * (Q * Q)
	exp := math.Exp(-expFract0 / expFract1) // https://ru.wikipedia.org/wiki/%D0%9D%D0%BE%D1%80%D0%BC%D0%B0%D0%BB%D1%8C%D0%BD%D0%BE%D0%B5_%D1%80%D0%B0%D1%81%D0%BF%D1%80%D0%B5%D0%B4%D0%B5%D0%BB%D0%B5%D0%BD%D0%B8%D0%B5
	fract := 1.0 / (Q * math.Sqrt(2.0*math.Pi))

	return fract * exp
}
func calcEstimatedNormProbabilities(Q float64, A float64, min float64, max float64, stepsAmount int) []float64 {
	rng := max - min
	stepLength := rng / float64(stepsAmount)

	probabilities := []float64{}

	for i := 0; i < stepsAmount; i++ {
		lft := min + float64(i)*stepLength
		rgt := lft + stepLength

		prob := normDistributionProbDensity(Q, A, rgt) * stepLength

		probabilities = append(probabilities, prob)
	}

	//maxProbId := misc.FindMaxFloat(probabilities)
	//l := float64(maxProbId) * stepLength
	//r := l + stepLength
	//fmt.Println("MinNumber: ", min, ". MostFrequentPart: ", l, "-", r, ". MaxNumber: ", max, ". MostFreqPartProb: ", probabilities[maxProbId])

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

func displayValues(values []float64, buffer *Buffer, posX float64, posY float64, scaleX float64, scaleY float64) {
	graphPos := new(Vector2).Set(posX, posY)
	graphScale := new(Vector2).Set(scaleX, scaleY)
	graph := new(Graph).Init(graphPos, graphScale, STYLE_BARS)
	graph.Draw(values, buffer)
}

func main() {
	generator := new(random.Generator).Init(0)
	prompter := new(Prompter).Init("\n", "Incorrect input!")
	buffer := new(Buffer).Init(24, 200)

	numberAmount := prompter.RequestInteger("Enter numbers amount: ")
	//numberAmount := 10
	exponent := prompter.RequestFloat("Enter exponent grater than 0: ")
	Q := prompter.RequestFloat("Enter Q  grater than 0: ")
	A := prompter.RequestFloat("Enter A: ")

	even := generateEvenDistribution(numberAmount, generator)
	estimatedEvenProbabilities := calcEstimatedEvenProbabilities(even[misc.FindMinFloat(even)], even[misc.FindMaxFloat(even)], 20)
	actualEvenProbabilities := countProbabilities(even, 20, false)
	evenPearsonCriterion := pearsonCriterion(estimatedEvenProbabilities, actualEvenProbabilities)
	//displayValues(estimatedEvenProbabilities, buffer, 0.0, 0.0, 0.3, 8.0)
	displayValues(actualEvenProbabilities, buffer, 0.0, 0.0, 0.3, 8.0)

	exp := generateExpDistribution(exponent, numberAmount, generator)
	estimatedExpProbabilities := calcEstimatedExpProbabilities(exp[misc.FindMinFloat(exp)], exp[misc.FindMaxFloat(exp)], 20)
	actualExpProbabilities := countProbabilities(exp, 20, false)
	expPearsonCriterion := pearsonCriterion(estimatedExpProbabilities, actualExpProbabilities)
	//displayValues(estimatedExpProbabilities, buffer, 0.0, 0.0, 0.3, 2.0)
	displayValues(actualExpProbabilities, buffer, 0.35, 0.0, 0.3, 2.0)

	//norm := []float64{}
	norm := generateNormDistribution(Q, A, numberAmount, generator)
	//for i := 0; i < numberAmount; i++ {
	//	norm[i] += 3.0 //Offset them so they can be displayed on my shitty graph :)
	//}
	estimatedNormProbabilities := calcEstimatedNormProbabilities(Q, A, norm[misc.FindMinFloat(norm)], norm[misc.FindMaxFloat(norm)], 20)
	actualNormProbabilities := countProbabilities(norm, 20, false)
	normPearsonCriterion := pearsonCriterion(estimatedNormProbabilities, actualNormProbabilities)
	//displayValues(estimatedNormProbabilities, buffer, 0.35, 0.0, 0.3, 4.0)
	displayValues(actualNormProbabilities, buffer, 0.7, 0.0, 0.3, 4.0)

	fmt.Println("=================================================================================================================================================================================================")
	buffer.Print()
	fmt.Println("=================================================================================================================================================================================================")
	fmt.Println()
	fmt.Println("Pearson criterion: ")
	fmt.Println("even: ", evenPearsonCriterion)
	fmt.Println("exp: ", expPearsonCriterion)
	fmt.Println("norm: ", normPearsonCriterion)
	buffer.Clear()
}
