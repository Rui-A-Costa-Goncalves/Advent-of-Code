package main

import (
	"fmt"
	"regexp"
	"os"
	"strings"
)

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
		positionRegex, errposition := regexp.Compile("(([a-z]|[A-Z])+) *=")
		leftRegex, errLeft := regexp.Compile("(([a-z]|[A-Z])+) *,")
		rightRegex, errRight :=regexp.Compile(", *(([a-z]|[A-Z])+)")
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
	currentPosition := "AAA"
	
	for {
		if (currentPosition == "ZZZ") {
			break
		}

		if i >= len(instructions) {
			i = 0
		}
		switch instructions[i] {
			case "L":
				currentPosition = positionMap[currentPosition].left
			case "R":
				currentPosition = positionMap[currentPosition].right
		}
		
		counter++
		i++
	}
	fmt.Println(counter)
}


type neighbours struct {
	left string
	right string
}
