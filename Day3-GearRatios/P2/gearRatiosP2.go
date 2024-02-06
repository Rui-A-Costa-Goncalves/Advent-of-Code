package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)


func digitsTotheRight (s string, index int) (string,int) {
	numberDigits := string(s[index])
	iterationIndex := index + 1
	for {
		if (iterationIndex >= len(s)) {
			break
		}
		_, err := strconv.Atoi(string(s[iterationIndex])) 
		if err != nil {
			break
		}
		numberDigits += string(s[iterationIndex])
		iterationIndex++
	}
	return numberDigits,len(numberDigits)
}

func digitsTotheLeft (s string, index int) string{
	numberDigits := ""
	iterationIndex := index -1
	for {
		if (iterationIndex < 0) {
			break
		}
		_, err := strconv.Atoi(string(s[iterationIndex])) 
		if err != nil {
			break
		}
		numberDigits = string(s[iterationIndex]) + numberDigits
		iterationIndex--
	}
	return numberDigits
}



func validateGear ( s []string,lineIndex int,columnIndex int) int {
	var partNumbers []string
	for i := lineIndex -1 ; i <= lineIndex +1 ; i++ {
		if i >= 0 && i < len(s) {
			startIndex := columnIndex - 1
			endIndex := columnIndex + 1

			if columnIndex == 0 {
				startIndex = 0
			} else if columnIndex == len(s[i])-1 {
				endIndex = columnIndex
			}
			_,err := strconv.Atoi(string(s[i][startIndex]))
			//Case where the 1st one is part of a number
			if err == nil {
				digits, len := digitsTotheRight(s[i],startIndex)
				partNumber := digitsTotheLeft(s[i],startIndex) + digits
				partNumbers =append(partNumbers,partNumber)
				if (startIndex + len < endIndex) {
					_,err := strconv.Atoi(string(s[i][endIndex]))
					if err == nil {
						partNumber, _ := digitsTotheRight(s[i],endIndex)
						partNumbers = append(partNumbers,partNumber)
					}
				}
			} else if startIndex +1 != endIndex {
					_,err := strconv.Atoi(string(s[i][startIndex+1]))
					// Case where the second one is the first number
					if err == nil {
						partNumber, _ := digitsTotheRight(s[i],startIndex+1)
						partNumbers =append(partNumbers,partNumber)
					} else{
						_,errEnd := strconv.Atoi(string(s[i][endIndex]))
						//Case where the third one is the first number
						if errEnd == nil {
							partNumber, _ := digitsTotheRight(s[i],endIndex)
							partNumbers =append(partNumbers,partNumber)
						}
					}
			} else {
				_,errEnd := strconv.Atoi(string(s[i][endIndex]))
				//Case where the third one is the first number
				if errEnd == nil {
					partNumber, _ := digitsTotheRight(s[i],endIndex)
					partNumbers =append(partNumbers,partNumber)
				}
			}
		}	
	}
	if len(partNumbers) == 2 {
		ele1,_ := strconv.Atoi(partNumbers[0]) 
		ele2,_ := strconv.Atoi(partNumbers[1])

		return ele1 * ele2
	} else {
		return 0
	}
}


func main () {
	if len(os.Args) < 2 {
		fmt.Println("Please insert a file input")
		return
	}
	fileInput := os.Args[1]

	readFile,err := os.ReadFile(fileInput)

	if err != nil {
		fmt.Println("Error reading file, got",err)
		return
	}
	data := string(readFile)

	dataLines := strings.Split(data, "\n")

	acc := 0

	for i,_ := range(dataLines) {
		charsInLine := strings.Split(dataLines[i], "")
		for j, char:= range (charsInLine) {
			if char == "*" {
				acc += validateGear(dataLines,i,j)
			}
		}	
	}
	fmt.Println(acc)
}
