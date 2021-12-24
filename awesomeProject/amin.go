package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//import "unicode/utf8"
//
////фнукция нужна чтоб составить алфавит букв.
//func removeDuplicateValues(intSlice []rune) []rune {
//	keys := make(map[rune]bool)
//	list := []rune{}
//
//	// If the key(values of the slice) is not equal
//	// to the already present value in new slice (list)
//	// then we append it. else we jump on another element.
//	for _, entry := range intSlice {
//		if _, value := keys[entry]; !value {
//			keys[entry] = true
//			list = append(list, entry)
//		}
//	}
//	return list
//}
//
////просто считает частотность букв. По частотности можно догадаться, на какие буквы была произведена замена.
//func calculateRunesFrequecy(encryptedText string) {
//	textLength := utf8.RuneCountInString(encryptedText)
//	fmt.Println("text length: ", textLength, " characters")
//
//	encryptedTextArray := []rune(encryptedText)
//	alph := removeDuplicateValues(encryptedTextArray) // строим алфавит
//
//	fmt.Println(alph)
//
//	for i := 0; i < len(alph); i++ { //берем каждую букву алфавита
//		count := 0
//		for j := 0; j < len(encryptedTextArray); j++ {
//			if alph[i] == encryptedTextArray[j] {
//				count++ // подсчитываем сколько раз встретилась буква в тексте.
//			}
//		}
//		//количество вхождений символа делим на длину текста. Получаем частоту в процентах.
//		fmt.Println(string(alph[i]), " ", float64(count)/float64(textLength))
//	}
//
//}
//
//// shows specified prompt and returns one line entered by user
//func enterTextPrompt(promptText string) string {
//	reader := bufio.NewReader(os.Stdin)
//
//	fmt.Println("")
//	fmt.Println("///////////////////////////////////////////////////////////////////")
//	fmt.Println(promptText)
//
//	text, _ := reader.ReadString('\n')
//	//convert CRLF to LF
//	text = strings.Replace(text, "\n", "", -1)
//
//	return text
//}
//
//// returns pointer to file from specified path
//func getFile(filePath string) (filePointer *os.File) {
//	filePointer, err := os.Open(filePath)
//	if err != nil {
//		fmt.Println("Failed to read file")
//		os.Exit(1)
//	}
//
//	return filePointer
//}
//
//func readTextFile(filePointer *os.File) string {
//	scanner := bufio.NewScanner(filePointer)
//
//	scanner.Split(bufio.ScanLines)
//	var text string
//
//	for scanner.Scan() {
//		text = text + "\n" + scanner.Text()
//	}
//
//	return text
//}
//
//func applyOffsets(encryptedTextArray []rune, keyword []rune) []rune {
//	for m := 1; m < len(encryptedTextArray); m++ {
//
//		if encryptedTextArray[m] == ' ' {
//			encryptedTextArray[m] = 'Я'
//		}
//
//		keyIndex := (m - 1) % len(keyword)
//		kernel := 'а' - 1 // the symbol right before russian 'a'
//		offset := keyword[keyIndex] - kernel
//
//		encryptedTextArray[m] -= offset
//
//		if encryptedTextArray[m] < kernel {
//			encryptedTextArray[m] += 33 //33 letters in russian alph
//		}
//
//		if encryptedTextArray[m] == 'Я' {
//			encryptedTextArray[m] = ' '
//		}
//
//	}
//
//	return encryptedTextArray
//}
//
//func findMostFrequentLetter(letters []rune) rune {
//	maxCount := 0
//	maxCountLetter := ' '
//
//	localAlph := removeDuplicateValues(letters) // make alph of existed letters (may reduce letters amount in some cases)
//
//	for k := 0; k < len(localAlph); k++ { // for each letter from alph
//		count := 0
//
//		for n := 0; n < len(letters); n++ { // count how much of them in selected letters array
//			if localAlph[k] == letters[n] {
//				count++
//			}
//		}
//
//		if count > maxCount {
//			maxCount = count
//			maxCountLetter = localAlph[k] //save letter if it appears more frequently
//		}
//	}
//
//	return maxCountLetter
//}
//
//func guessKeyWord(encryptedTextArray []rune, keyWordLength int) []rune {
//	keyword := []rune{}
//
//	for i := 1; i < keyWordLength+1; i++ {
//		letters := []rune{}
//
//		for j := i; j < len(encryptedTextArray); j += keyWordLength {
//			if j < len(encryptedTextArray) {
//				letters = append(letters, encryptedTextArray[j]) // take each letter with step length equals keyWordLength
//			}
//		}
//
//		maxCountLetter := findMostFrequentLetter(letters)
//
//		keyword = append(keyword, maxCountLetter) //after checking all the letters the most frequent one goes to the keyword
//	}
//
//	return keyword
//}
//
//func main() {
//	/////////////////////////криптопротоколы, первая лаба.
//	//filePath := enterTextPrompt("Enter path to file with encrypted text.")
//	//file := getFile(filePath)
//	//encryptedText := readTextFile(file)
//	//
//	//calculateRunesFrequecy(encryptedText)
//
//	////////////////////////криптопротоколы, вторая лаба
//	filePath := enterTextPrompt("Enter path to file with encrypted text.")
//	file := getFile(filePath)
//
//	keyWordLengthText := enterTextPrompt("Enter estimated keyword length.")
//	keyWordLength, _ := strconv.Atoi(keyWordLengthText)
//
//	encryptedText := readTextFile(file)
//	encryptedTextArray := []rune(encryptedText)
//
//	fmt.Println("Presumably a keyword is: ")
//	keyword := guessKeyWord(encryptedTextArray, keyWordLength)
//	fmt.Println(string(keyword))
//
//	fmt.Println("Decoded text.")
//	encryptedTextArray = applyOffsets(encryptedTextArray, keyword)
//	fmt.Println(string(encryptedTextArray))
//}
