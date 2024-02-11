package main

import (
	"fmt"
	"strings"
	"os"
	"regexp"
	"strconv"
)

func isAllZeros (integers []int) bool {
	for i := range(integers) {
		if integers[i] != 0 {
			return false
		}
	}
	return true
}


func main () {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a input file")
		return
	}
	fileName := os.Args[1]

	file,errFile := os.ReadFile(fileName)
	if errFile != nil {
		fmt.Println("Error reading file, got",errFile)
		return
	}
	numbersRegex, errNumber := regexp.Compile("-?[0-9]+")
	if errNumber != nil {
		fmt.Println("Error defining the regex, got",errNumber)
	}

	data := string(file)
	dataLines := strings.Split(data,"\n")

	var sequences[][]int

	for _,line := range(dataLines) {
		numbersString := numbersRegex.FindAllString(line,-1)
		var numbers []int
		for _, ns := range(numbersString) {
			number, errNumber := strconv.Atoi(ns)
			if errNumber != nil {
				fmt.Println("Got error converting to number, got", errNumber)
			}
			numbers = append(numbers, number)
		}
		sequences = append(sequences, numbers)
	}
	sum := 0
	for _, sequence := range(sequences) {
		var firstElems []int
		actualSequence := sequence
		firstElems = append(firstElems, sequence[0])
		for !isAllZeros(actualSequence){
			var newSequence []int
			for i:=0; i < len(actualSequence)-1; i++ {
				diff := actualSequence[i +1] - actualSequence[i]
				newSequence = append(newSequence, diff)
			}
			actualSequence = newSequence
			firstElems = append(firstElems, actualSequence[0])
		}
		sequenceValue := 0
		for i:= len(firstElems)-1; i >= 0; i-- {
			sequenceValue = firstElems[i]-sequenceValue
		}
		sum += sequenceValue
	}
	fmt.Println(sum)
}