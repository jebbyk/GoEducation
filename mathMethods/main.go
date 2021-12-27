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
	"time"
)

func countProbabilities(values []float64, steps int) []float64 {
	min := misc.FindMinFloat(values)
	max := misc.FindMaxFloat(values)

	rng := max - min //TODO take negative number into account later

	stepLength := rng / float64(steps)

	probabilities := []float64{}

	for i := 0; i < steps; i++ {
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

func generateExpDistribution(exponent float64, iterations int, generator *random.Generator) []float64 {
	var values []float64
	for i := 0; i < iterations; i++ {
		value := -exponent * math.Log10(generator.Fgen())
		values = append(values, value)
	}

	return values
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

	//iterations := prompter.RequestInteger("Enter iterations amount: ")
	//iterations := 10
	exponent := prompter.RequestFloat("Enter exponent grater than 0: ")
	Q := prompter.RequestFloat("Enter Q between grater than 0: ")

	norm := []float64{}

	for iterations := 10; iterations < 100000; iterations += 2 { // add 2 in each iteration cuz normDistribution generates 2 numbers at once
		even := generateEvenDistribution(iterations, generator)
		displayValues(even, buffer, 0.0, 0.0, 0.3, 8.0, 100)

		exp := generateExpDistribution(exponent, iterations, generator)
		displayValues(exp, buffer, 0.35, 0.0, 0.3, 8.0, 100)

		norm = generateNormDistribution(Q, iterations, generator)
		for i := 0; i < iterations; i++ {
			norm[i] += 2.0 //Offset them so they can be displayed on my shitty graph :)
		}
		displayValues(norm, buffer, 0.7, 0.0, 0.3, 8.0, 100)

		//check := []float64{
		//	1.4189092916253032,
		//	-1.4249173176498258,
		//	1.466958001903065,
		//	-0.829295185885189,
		//	-1.3796978679412768,
		//	0.6629034680790691,
		//	-0.5765531909639737,
		//	0.25304803914971186,
		//	0.004704684204398089,
		//	-0.6770376932574126,
		//	0.4168847612405232,
		//	0.7347082010528398,
		//	-1.227555460752853,
		//	-1.5862556899372942,
		//	1.0071168644174406,
		//	0.29785049246314366,
		//	-0.27062178802202214,
		//	0.13024576379502473,
		//	-0.5194136414193193,
		//	-0.12841639819991565,
		//	-0.015244704127991216,
		//	0.3100950895055931,
		//	0.7510028734027906,
		//	-1.268510270849423,
		//	-2.80346608765641,
		//	-0.24822144294931728,
		//	-0.21039672180040836,
		//	0.8975922260697201,
		//	-1.0489269948134343,
		//	-0.6770096935866017,
		//	-0.6018820793856686,
		//	1.003833414251205,
		//	0.10185820203667705,
		//	-1.7710007982059348,
		//	0.17194554460711045,
		//	-0.5539392532892405,
		//	-0.15994524087812517,
		//	0.6680810909953735,
		//	-0.526425996199556,
		//	0.5519548565216483,
		//	-1.4059468284447623,
		//	-1.4401868220711336,
		//	2.119157240498979,
		//	-0.8334827635603873,
		//	-0.17727855508345008,
		//	0.7643923839300037,
		//	-0.4693361444055294,
		//	-1.5448025561386995,
		//	-1.9858945032318978,
		//	-0.06361229268063669,
		//	1.261383696234355,
		//	1.336196794380638,
		//	0.4204142282426973,
		//	0.6245166452019515,
		//	-0.453841554029363,
		//	1.453236562565096,
		//	0.29035741703752127,
		//	-0.1431342415603266,
		//	-0.845840830934785,
		//	-1.6855302215411578,
		//	0.252442031054777,
		//	-0.33413390855220326,
		//	-1.613822688779097,
		//	0.43833516520054333,
		//	0.6704888618113668,
		//	-0.7285944857803375,
		//	1.6356636581319615,
		//	1.4417710324523667,
		//	-0.9171339881610582,
		//	0.5423273040757846,
		//	1.8596260898308834,
		//	-0.9882126330874942,
		//	1.600791527917314,
		//	-0.293358433667752,
		//	-0.5352744460376923,
		//	0.573415878054984,
		//	0.5246408759827482,
		//	-0.5476654272021354,
		//	-0.6407798346198866,
		//	1.1238710786654669,
		//	-2.1225794194887913,
		//	-0.24182329519347093,
		//	-1.0987701237528582,
		//	1.25757653989811,
		//	1.6336174080266728,
		//	-0.03820007499939612,
		//	0.41706834268717885,
		//	-0.9138409565652144,
		//	-0.8909240984578742,
		//	-0.7255756757172187,
		//	0.35553576820073085,
		//	-1.9861753151273795,
		//	-0.7303380923052256,
		//	0.0682711580325713,
		//	0.740105417494352,
		//	1.3613794129180492,
		//	0.8328385446254055,
		//	-2.014630315573572,
		//	-0.7912245423784181,
		//	1.0406032335570308,
		//}
		//for i := 0; i < iterations; i++ {
		//	check[i] += 0.5
		//}
		//displayValues(check, buffer, 0.7, 0.0, 0.3, 2.0, 30)

		fmt.Println("=========================================================================================================================================================================================================")
		buffer.Print()
		fmt.Println("=========================================================================================================================================================================================================")

		buffer.Clear()

		time.Sleep(16 * time.Millisecond)
	}
}
