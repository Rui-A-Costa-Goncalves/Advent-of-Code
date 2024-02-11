package main

import (
	"fmt"
	"regexp"
	"os"
	"strings"
)

func GCD(a, b int) int {
	for b != 0 {
			t := b
			b = a % b
			a = t
	}
	return a
}


func MyLCM(integers []int) int {
	if len(integers) == 0 {
		return 0
	} 
	if len(integers) == 1 {
		return integers[0]
	}
	if len(integers) == 2 {
		return integers[0] * integers[1] / GCD(integers[0], integers[1])
	}

	result := integers[0] * integers[1] / GCD(integers[0], integers[1])
	var newInput []int
	newInput = append(newInput, result)
	newInput = append(newInput, integers[2:]...)

	return MyLCM(newInput)
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("File was not provided")
	}

	fileName := os.Args[1]
	file, errFile := os.ReadFile(fileName)
	if errFile != nil {
		fmt.Println("Error reading file, got", errFile)
	}
	data := string(file)

	dataLines := strings.Split(data,"\n")
	positionMap := make(map[string]neighbours)
	var instructions []string
	for i, line := range(dataLines) {
		if i == 0 {
			instructions = strings.Split(line,"")
		}
		positionRegex, errposition := regexp.Compile("(([0-9]|[a-z]|[A-Z])+) *=")
		leftRegex, errLeft := regexp.Compile("(([0-9]|[a-z]|[A-Z])+) *,")
		rightRegex, errRight :=regexp.Compile(", *(([0-9]|[a-z]|[A-Z])+)")
		if errposition != nil || errLeft != nil || errRight != nil {
			fmt.Println("Error defining one of the regex")
		}
		if positionRegex.MatchString(line) {
			matches := positionRegex.FindStringSubmatch(line)
			position := matches[1]

			leftMatches := leftRegex.FindStringSubmatch(line)
			left := leftMatches[1]

			rightMatches := rightRegex.FindStringSubmatch(line)
			
			right := rightMatches[1]

			neighbours := neighbours{left : left, right : right}
			positionMap[position] = neighbours
		}
	}

	counter := 0
	i:= 0
	var currentPositions []string
	for position, _ := range(positionMap) {
		if strings.Split(position,"")[len(position)-1] == "A" {
			currentPositions = append(currentPositions, position)
		}
	}
	terminationSteps := make(map[int]int)
	
	for {
		if (len(terminationSteps) == len(currentPositions)) {
			break
		}
		if i >= len(instructions) {
			i = 0
		}
		for j := 0 ; j < len(currentPositions); j++ {
			//if it ends in Z
			if strings.Split(currentPositions[j],"")[len(currentPositions[j])-1] == "Z" {
				_, ok := terminationSteps[j]
				if ok == false {
					terminationSteps[j] = counter
				}
			}
			switch instructions[i] {
				case "L":
					currentPositions[j] = positionMap[currentPositions[j]].left
				case "R":
					currentPositions[j] = positionMap[currentPositions[j]].right
			}
		} 
		counter++
		i++
	}
	var myList []int
	for _,val := range(terminationSteps) {
		myList = append(myList, val)
	}
	
	fmt.Println(MyLCM(myList))
}


type neighbours struct {
	left string
	right string
}
