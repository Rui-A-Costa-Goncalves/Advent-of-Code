package main

import (
	"fmt"
	"os"
	"strings"
	"strconv"
)

func validate ( s []string,lineIndex int,initialj int,finalj int) bool {
	for i := lineIndex -1 ; i <= lineIndex +1 ; i++ {
		if i >= 0 && i < len(s) {
			for j := initialj -1; j <= finalj +1; j++ {
				if j >= 0 && j < len(s[i]) {
					if string(s[i][j]) != "." && (i != lineIndex || (j < initialj || j > finalj)) {
						return true
					}
				}
			}
		}
		
	}
	return false
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
		initialj := -1
		finalj := -1
		number := ""
		for j, char:= range (charsInLine) {
			_,errConv := strconv.Atoi(char)
			if errConv == nil {
				if initialj == -1 {
					initialj = j
				}
				number += char
				finalj = j
				if j == len(charsInLine)-1 {
					if validate(dataLines,i,initialj,finalj) {
						numberInt, _ := strconv.Atoi(number)
						acc += numberInt
					}
				}

			} else {
				if number != "" && validate(dataLines,i,initialj,finalj) {
					numberInt, _ := strconv.Atoi(number)
					acc += numberInt
				}
				initialj = -1
				finalj = -1
				number = ""
			}
		}
	}
	fmt.Println(acc)

}