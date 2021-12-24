package prompt

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Prompter struct {
	delimeter    string
	errorMessage string
}

func (p *Prompter) Init(delimiter string, errorMessage string) *Prompter {
	p.delimeter = delimiter
	p.errorMessage = errorMessage
	return p
}

// shows specified prompt and returns one line entered by user
func (p *Prompter) RequestString(promptText string) string {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println(p.delimeter)
	fmt.Println(promptText)

	text, _ := reader.ReadString('\n')
	//convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	return text
}

func (p *Prompter) RequestInteger(promptText string) int {
	text := p.RequestString(promptText)
	value, err := strconv.Atoi(text)
	if err != nil {
		fmt.Println(p.errorMessage)
		value = p.RequestInteger(promptText) // will ask for correct input recursively
	}
	return value
}
