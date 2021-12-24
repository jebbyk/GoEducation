package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// shows specified prompt and returns one line entered by user
func enterTextPrompt(promptText string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("///////////////////////////////////////////////////////////////////")
	fmt.Println(promptText)

	text, _ := reader.ReadString('\n')
	//convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func main() {
	//kernel := 151875
	//multiplier := 31
	//delta := 1283
	//x := 17
	//
	//stepsAmountText := enterTextPrompt("Enter key.")
	//stepsAmount, _ := strconv.Atoi(stepsAmountText)
	//
	//showIntermediateResults := enterTextPrompt("Display intermediate results: [y/n]")
	//
	//if showIntermediateResults == "y" {
	//	for i := 0; i < stepsAmount; i++ {
	//		x = (multiplier*x + delta) % kernel
	//		fmt.Println("№: ", i, ". Sample: ", x)
	//	}
	//} else {
	//
	//	enableOptimization := enterTextPrompt("Apply cycle optimization: [y/n]")
	//
	//	if enableOptimization == "y" {
	//		stepsAmount = stepsAmount % kernel
	//	}
	//
	//	startTime := time.Now()
	//
	//	for i := 0; i < stepsAmount; i++ {
	//		x = (multiplier*x + delta) % kernel
	//	}
	//
	//	endTime := time.Now()
	//
	//	fmt.Println("Result: ", x, ". Time spent: ", endTime.Sub(startTime))
	//}

	generator := new(Generator)

}
