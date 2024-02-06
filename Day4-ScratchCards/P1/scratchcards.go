package main

import (
	"fmt"
	"os"
	"strings"
	"regexp"
	"math"
)

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

	acc := 0

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

		occurences := 0


		//implementing a set
		/*
		uniqueWinningNumbers := map[string]struct{}{}
		for _,winningNumber := range(myNumbers) {
			_,ok := uniqueWinningNumbers[winningNumber]
			if ok == false {
				uniqueWinningNumbers[winningNumber] = struct {}{}
			}
		}*/

		for i := range (myNumbers) {
			for j := range(winningNumbers) {
				if myNumbers[i] == winningNumbers[j] {
					occurences++
				}
			}
		}
		
		acc += int(math.Pow(2,float64(occurences-1)))
	}
	fmt.Println(acc)
}