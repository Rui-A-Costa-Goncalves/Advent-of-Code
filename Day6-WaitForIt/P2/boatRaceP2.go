package main

import (
	"fmt"
	"regexp"
	"os"
	"strings"
	"strconv"
)

func getNumber (line string)int {
	
	numbersRegex, errRegex := regexp.Compile("[0-9]+")
	if errRegex != nil {
		fmt.Println("Error defining the regex, got",errRegex)
	}
	numbersString := numbersRegex.FindAllString(line, -1)
	var totalNumber string
	
	for _,number := range(numbersString) {
		totalNumber += number
	}
	number, errInt := strconv.Atoi(totalNumber)
	if errInt != nil {
		fmt.Println("Error converting string to int, got",errInt)
	}
	
	return number
}

func main () {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a file")
		return
	}
	fileName := os.Args[1]
	file, errFile := os.ReadFile(fileName)
	if errFile != nil {
		fmt.Println("Error reading file, got",errFile)
		return
	}
	data := string(file)
	dataLines := strings.Split(data, "\n")
	time := getNumber(dataLines[0])
	distance := getNumber(dataLines[1])

	var minNumber int	
	for j := 1; j< time; j++ {
		if (time - j)*j > distance {
			minNumber = j
			break
		} 
	}

	numberOfWays := time - minNumber*2 +1
	
	fmt.Println(numberOfWays)
}