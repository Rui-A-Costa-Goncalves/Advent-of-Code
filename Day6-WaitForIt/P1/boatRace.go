package main

import (
	"fmt"
	"regexp"
	"os"
	"strings"
	"strconv"
)

func getNumbers (line string) []int {
	
	numbersRegex, errRegex := regexp.Compile("[0-9]+")
	if errRegex != nil {
		fmt.Println("Error defining the regex, got",errRegex)
	}
	numbersString := numbersRegex.FindAllString(line, -1)
	var numbers []int
	
	for _,number := range(numbersString) {
		numberInt, errInt := strconv.Atoi(number)
		if errInt != nil {
			fmt.Println("Error converting string to int, got",errInt)
		} else {
			numbers = append(numbers,numberInt)
		}
	}
	return numbers
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
	times := getNumbers(dataLines[0])
	distances := getNumbers(dataLines[1])

	var minNumbers []int
	for i := range(times) {
		for j := 1; j<times[i]; j++ {
			if (times[i] - j)*j > distances[i] {
				minNumbers = append(minNumbers,j)
				break
			} 
		}
	}

	var numberOfWays []int
	for i := range(times) {
		numberOfWays = append(numberOfWays, times[i] - minNumbers[i]*2 +1)
	}

	totalNumberOfWays := 1
	for _,number := range(numberOfWays) {
		totalNumberOfWays *= number
	}
	fmt.Println(totalNumberOfWays)
}