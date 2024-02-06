package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"strconv"
	"sort"
)

type card struct {
	hand []string
	bid int
}

func value (char string) int {
	switch char {
	case "A":
		return 13
	case "K":
		return 12
	case "Q":
		return 11
	case "J":
		return 10
	case "T":
		return 9
	case "9":
		return 8
	case "8":
		return 7
	case "7":
		return 6
	case "6":
		return 5
	case "5":
		return 4
	case "4":
		return 3
	case "3":
		return 2
	default:
		return 1
	}
}

func numberOfEqual (hand []string) map[string]int {
	ocurrences := make(map[string]int)
	for _,card := range(hand) {
		acc := 0
		for i := range(hand) {
			if hand[i] == card {
				acc ++
			} 
		}
		_,ok := ocurrences[card]
		if ok == false {
			ocurrences[card] = acc
		}
	}
	return ocurrences
}

func typeOfHand (c card) int {
	ocurrences := numberOfEqual(c.hand)
	if len(ocurrences) == 1 {
		return 7
	}

	for key,val := range(ocurrences) {
		if val == 4 {
			return 6
		}
		if val == 3 {
			for _,newVal := range ocurrences {
				if newVal == 2 {
					return 5
				}
			}
			return 4
		}
		if val == 2 {
			for newKey,newVal := range ocurrences {
				if newVal == 2 && newKey != key {
					return 3
				}
				if newVal == 3 {
					return 5
				}
			}
			return 2
		}
	}
	return 1
}

func isValueSmaller(a,b card) bool {
	if typeOfHand(a) != typeOfHand(b) {
		return typeOfHand(a) < typeOfHand(b)
	} else {
		for i,char := range a.hand {
			if char != b.hand[i] {
				return value(char) < value(b.hand[i])
			}
		}
		return true
	}
}


func main () {
	if len(os.Args) < 2 {
		fmt.Println("No file was inserted")
		return
	}

	fileName := os.Args[1]

	file,errFile := os.ReadFile(fileName)
	if errFile != nil {
		fmt.Println("Error reading file, got",errFile)
		return
	}

	data := string(file)
	dataLines := strings.Split(data,"\n")

	var cards []card

	for _,line := range(dataLines) {
		parts := strings.Split(line, " ")
		charRegex, errChar := regexp.Compile("[a-z]|[A-Z]|[0-9]")
		if errChar != nil {
			fmt.Println("Error defining the char regex, got",errChar)
		}
		hand := charRegex.FindAllString(parts[0],-1)

		numberRegex, errNum := regexp.Compile("[0-9]+")
		if errNum != nil {
			fmt.Println("Error deining the number regex, got", errNum)
		}
		value := numberRegex.FindString(parts[1])
		valueInt,errInt := strconv.Atoi(value)
		if errInt != nil {
			fmt.Println("Error converting the value to int, got", errInt)
		}
		card := card{hand : hand, bid : valueInt}
		cards = append(cards,card)
	}

	sort.Slice(cards, func(i,j int) bool {
		return isValueSmaller(cards[i],cards[j])
	})

	result := 0
	for i := range(cards) {
		value := cards[i].bid * (i+1)
		result += value
	}

	fmt.Println(result)

}