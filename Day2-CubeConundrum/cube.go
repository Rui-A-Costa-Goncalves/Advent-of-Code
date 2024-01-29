package main

import (
	"fmt"
	"os"
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func doPart1 (scanner bufio.Scanner) {
	var acc int
	for scanner.Scan() {
		idReg,err := regexp.Compile("Game (.+):")
		if err != nil {
			fmt.Println("Error on the regex, got ",err)
			return
		}

		idString := idReg.FindStringSubmatch(scanner.Text())[1]


		id,errAtoi := strconv.Atoi(idString)
		if errAtoi!= nil {
			fmt.Println("Error when converting id to int, got ",errAtoi)
			return
		}

		games := strings.Split(scanner.Text(),";")

		redFinder,errRed := regexp.Compile("([0-9]*) red")
		greenFinder,errGreen := regexp.Compile("([0-9]*) green")
		blueFinder,errBlue := regexp.Compile("([0-9]*) blue")

		if (errRed != nil || errGreen != nil || errBlue != nil) {
			fmt.Println("Error on the color regex")
			return
		}

		var maxGreen, maxRed, maxBlue int
		fmt.Println(games)
		for _,game := range(games) {
			reds := redFinder.FindStringSubmatch(game)
			if reds != nil {
				numberOfReds,errConv := strconv.Atoi(reds[1])
				if (errConv != nil) {
					fmt.Println("Error converting number of reds")
					return
				}
				if maxRed < numberOfReds {
					maxRed = numberOfReds
				} 
			}
			greens := greenFinder.FindStringSubmatch(game)
			if greens != nil {
				numberOfGreens,errConv := strconv.Atoi(greens[1])
				if (errConv != nil) {
					fmt.Println("Error converting number of greens")
					return
				}
				if maxGreen < numberOfGreens {
					maxGreen = numberOfGreens
				} 
			}
			blues := blueFinder.FindStringSubmatch(game)
			if blues != nil {
				numberOfBlues,errConv := strconv.Atoi(blues[1])
				if (errConv != nil) {
					fmt.Println("Error converting number of blues")
					return
				}
				if maxBlue < numberOfBlues {
					maxBlue = numberOfBlues
				} 
			}

		}
		if maxRed <= 12 && maxGreen <= 13 && maxBlue <= 14 {
			acc += id
		}
	}
	fmt.Println(acc)
}

func doPart2 (scanner bufio.Scanner) {
	var acc int
	for scanner.Scan() {
		games := strings.Split(scanner.Text(),";")

		redFinder,errRed := regexp.Compile("([0-9]*) red")
		greenFinder,errGreen := regexp.Compile("([0-9]*) green")
		blueFinder,errBlue := regexp.Compile("([0-9]*) blue")

		if (errRed != nil || errGreen != nil || errBlue != nil) {
			fmt.Println("Error on the color regex")
			return
		}

		var maxGreen, maxRed, maxBlue int
		fmt.Println(games)
		for _,game := range(games) {
			reds := redFinder.FindStringSubmatch(game)
			if reds != nil {
				numberOfReds,errConv := strconv.Atoi(reds[1])
				if (errConv != nil) {
					fmt.Println("Error converting number of reds")
					return
				}
				if maxRed < numberOfReds {
					maxRed = numberOfReds
				} 
			}
			greens := greenFinder.FindStringSubmatch(game)
			if greens != nil {
				numberOfGreens,errConv := strconv.Atoi(greens[1])
				if (errConv != nil) {
					fmt.Println("Error converting number of greens")
					return
				}
				if maxGreen < numberOfGreens {
					maxGreen = numberOfGreens
				} 
			}
			blues := blueFinder.FindStringSubmatch(game)
			if blues != nil {
				numberOfBlues,errConv := strconv.Atoi(blues[1])
				if (errConv != nil) {
					fmt.Println("Error converting number of blues")
					return
				}
				if maxBlue < numberOfBlues {
					maxBlue = numberOfBlues
				} 
			}

		}
		acc += maxRed*maxBlue*maxGreen
	}
	fmt.Println(acc)
}



func main() {

	if len(os.Args)< 3 {
		fmt.Println("Please provide a file as input and the part that you want.")
		return
	}

	fileName := os.Args[1]

	file,err := os.Open(fileName)

	if err != nil {
		fmt.Println("Error opening file, got the error ",err)
		return
	}

	scanner := bufio.NewScanner(file)

	part := os.Args[2]

	switch part {
	case "2":
		doPart2(*scanner)

	default:
		doPart1(*scanner)
	}


	

}