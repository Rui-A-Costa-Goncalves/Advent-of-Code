package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
)

func numberOfCards(index int, cards [][]int) int{
	if len(cards[index]) == 0 {
		return 1
	}
	acc := 0
	for j:= range cards[index] {
		acc += numberOfCards(index+j+1,cards)
	}
	return 1 + acc
}


func main () {
	if (len(os.Args) < 2) {
		fmt.Println("Please insert the file")
		return
	}

	filename := os.Args[1]

	file, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error reading file, got",err)
		return
	}

	data := string(file)
	dataLines := strings.Split(data, "\n")

	var generatedCards [][]int
	
	for _,line := range(dataLines) {
		cards := strings.Split(line,"|")

		winningRegex, errWinning := regexp.Compile("Card +[0-9]+:( *([0-9]+ *)+)")
		if errWinning != nil {
			fmt.Println("Error compiling regex, got",errWinning)
			return
		}
		numbersRegex, errnumbers := regexp.Compile("[0-9]+")
		if errnumbers != nil {
			fmt.Println("Error compiling regex, got",errnumbers)
			return
		}
		matches := winningRegex.FindStringSubmatch(cards[0])
		match := matches[1]

		winningNumbers := numbersRegex.FindAllString(match,-1)
		myNumbers := numbersRegex.FindAllString(cards[1],-1)

		var cardsI []int
		for i := range (myNumbers) {
			for j := range(winningNumbers) {
				if myNumbers[i] == winningNumbers[j] {
					cardsI = append(cardsI,1)
				}
			}
		}
		generatedCards = append(generatedCards,cardsI)
	}
	acc := 0
	for i := range (dataLines) {
		acc += numberOfCards(i,generatedCards)
	}

	fmt.Println(acc)
}