package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

var translator = map[string]string{
	"one":"1",
	"two":"2",
	"three":"3",
	"four":"4",
	"five":"5",
	"six":"6",
	"seven":"7",
	"eight":"8",
	"nine":"9",
}

func LastCharachter (s string) string {
	newS := strings.Split(s,"")
	index := len(newS) -1
	return newS[index]
}


func replaceDigit (s string) string {
	for key,value := range(translator) {
		if strings.HasPrefix(s,key) {
			newString := strings.Replace(s,key,value,1)
			return newString[:1] + LastCharachter(key) + newString[1:]
		}
	}
	return s
	
}


func formatString(s string) string{
	newString := s
	index := 0

	for {
		if (index >= len(newString)) {
			break
		}
		aux := newString[index:]
		answerString := replaceDigit(aux)
		if answerString != aux {
			newString = strings.Replace(newString,newString[index:],answerString,1)
		}
		index ++
	}
	return newString
}



func main () {

	if len(os.Args) < 2 {
		fmt.Println("Please provide a file as input")
		return
	}
	fileName:= os.Args[1]

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("File was not read, the error was",err)
		return
	}

	defer file.Close() //Closes at the end of the execution
	scanner := bufio.NewScanner(file)

	var acc int
	

	//Scan works by line, if it hits a EOF it will exit out of the for loop
	for scanner.Scan() {
		first := ""
		last := ""
		// The default split ("") is by character
		newString := formatString(scanner.Text())
		for _, char := range strings.Split(newString, "") {
			_,errInt := strconv.Atoi(char) 
			
			if (errInt == nil) {
				if first == "" {
					first = char
				}
				last = char
				
			}
		}
		toSumStr := first + last
		toSum, _ := strconv.Atoi(toSumStr)
		fmt.Println(toSum)

		acc += toSum
	}
	fmt.Println(acc)
}

