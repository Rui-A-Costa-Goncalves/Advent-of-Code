package main

import (
	"fmt"
	"os"
	"strings"
	"errors"
	"regexp"
	"strconv"
)

func getNumbers (line string,numbersRegex regexp.Regexp) []int {
	numbers := numbersRegex.FindAllString(line,-1)
	var numbersInt []int
	for _,number := range (numbers) {
		numberInt,errInt := strconv.Atoi(number)
		if (errInt != nil) {
			fmt.Println("Error converting to int, got",errInt)
		}
		numbersInt = append(numbersInt,numberInt)
	}
	return numbersInt
}

func getNextValue (phase string,mp map[string][]sourceMapper, numbertoSearch int) int {
	result :=-1
	for _,sm := range(mp[phase]) {
		res,err := getNext(sm,numbertoSearch)
		if err == nil {
			result= res
		}
	}
	if result == -1 {
		result = numbertoSearch
	}

	switch phase {
		case "soil":
			return getNextValue("fertilizer",mp,result)
		case "fertilizer":
			return getNextValue("water",mp,result)
		case "water":
			return getNextValue("light",mp,result)
		case "light":
			return getNextValue("temperature",mp,result)
		case "temperature":
			return getNextValue("humidity",mp,result)
		case "humidity":
			return getNextValue("location",mp,result)
		default:
			return result
	}
}

func getNext(sm sourceMapper,numbertoSearch int) (int,error) {
	if (numbertoSearch >= sm.sourceBegin && numbertoSearch <= sm.sourceEnd) {
		return sm.destinyBegin + numbertoSearch-sm.sourceBegin,nil
	}

	return 0, errors.New("Index out of bounds")
}

func createSourceMapper (destinyBegin int, sourceBegin int, length int) sourceMapper {
	newSourceEnd := sourceBegin + length -1
	return sourceMapper{sourceBegin: sourceBegin, sourceEnd: newSourceEnd, destinyBegin: destinyBegin}
}


func main() {
	if (len(os.Args)< 2) {
		fmt.Println("No file provided")
		return
	}
	fileName := os.Args[1]
	file,errFile := os.ReadFile(fileName)
	if errFile != nil {
		fmt.Println("Got error finding the gile, got",errFile)
		return
	}


	data := string(file)

	dataLines := strings.Split(data, "\n")
	
	var seeds []int
	context := "seed"


	seedsRegex,errSeed := regexp.Compile("seeds:( *([0-9]+ *)+)")
	numbersRegex, errnumbers := regexp.Compile("[0-9]+")
	soilRegex,errSoil := regexp.Compile("seed-to-soil map:")
	fertRegex,errFert := regexp.Compile("soil-to-fertilizer map:")
	waterRegex,errWater := regexp.Compile("fertilizer-to-water map:")
	lightRegex,errLight := regexp.Compile("water-to-light map:")
	tempRegex,errTemp := regexp.Compile("light-to-temperature map:")
	humRegex,errHum := regexp.Compile("temperature-to-humidity map:")
	locationRegex,errLocation := regexp.Compile("humidity-to-location map:")
	if (errSeed != nil || errSoil != nil || errFert != nil|| errWater != nil|| errLight != nil|| errTemp != nil || errHum != nil || errnumbers != nil || errLocation != nil) {
		fmt.Println("Error defining the regex")
	}


	seedMap := make(map[string][]sourceMapper)

	for _,line := range(dataLines) {
		switch context {
			case "seed":
				if seedsRegex.MatchString(line) {
					matches := seedsRegex.FindStringSubmatch(line)
					match := matches[1]
					seeds = getNumbers(match,*numbersRegex)
					
					
				} else if soilRegex.MatchString(line) {
					context = "soil"
				}
			case "soil":
				if numbersRegex.MatchString(line) {
					numbersSoil := getNumbers(line,*numbersRegex)
					newSourceMapper := createSourceMapper(numbersSoil[0],numbersSoil[1],numbersSoil[2])
					_,ok := seedMap[context]
					if (ok == false) {
						var list []sourceMapper
						list = append(list,newSourceMapper)
						seedMap[context] = list
					} else {
						seedMap[context] = append(seedMap[context],newSourceMapper)
					}

				} else if fertRegex.MatchString(line){
					context = "fertilizer"
				}
			case "fertilizer":
				if numbersRegex.MatchString(line) {
					numbersFert := getNumbers(line,*numbersRegex)
					newSourceMapper := createSourceMapper(numbersFert[0],numbersFert[1],numbersFert[2])
					_,ok := seedMap[context]
					if (ok == false) {
						var list []sourceMapper
						list = append(list,newSourceMapper)
						seedMap[context] = list
					} else {
						seedMap[context] = append(seedMap[context],newSourceMapper)
					}

				} else if waterRegex.MatchString(line){
					context = "water"
				}
			case "water":
				if numbersRegex.MatchString(line) {
					numbersWater := getNumbers(line,*numbersRegex)
					newSourceMapper := createSourceMapper(numbersWater[0],numbersWater[1],numbersWater[2])
					_,ok := seedMap[context]
					if (ok == false) {
						var list []sourceMapper
						list = append(list,newSourceMapper)
						seedMap[context] = list
					} else {
						seedMap[context] = append(seedMap[context],newSourceMapper)
					}

				} else if lightRegex.MatchString(line){
					context = "light"
				}
			case "light":
				if numbersRegex.MatchString(line) {
					numbersLight := getNumbers(line,*numbersRegex)
					newSourceMapper := createSourceMapper(numbersLight[0],numbersLight[1],numbersLight[2])
					_,ok := seedMap[context]
					if (ok == false) {
						var list []sourceMapper
						list = append(list,newSourceMapper)
						seedMap[context] = list
					} else {
						seedMap[context] = append(seedMap[context],newSourceMapper)
					}

				} else if tempRegex.MatchString(line){
					context = "temperature"
				}
			case "temperature":
				if numbersRegex.MatchString(line) {
					numbersTemp := getNumbers(line,*numbersRegex)
					newSourceMapper := createSourceMapper(numbersTemp[0],numbersTemp[1],numbersTemp[2])
					_,ok := seedMap[context]
					if (ok == false) {
						var list []sourceMapper
						list = append(list,newSourceMapper)
						seedMap[context] = list
					} else {
						seedMap[context] = append(seedMap[context],newSourceMapper)
					}

				} else if humRegex.MatchString(line){
					context = "humidity"
				}

			case "humidity":
				if numbersRegex.MatchString(line) {
					numbersHum := getNumbers(line,*numbersRegex)
					newSourceMapper := createSourceMapper(numbersHum[0],numbersHum[1],numbersHum[2])
					_,ok := seedMap[context]
					if (ok == false) {
						var list []sourceMapper
						list = append(list,newSourceMapper)
						seedMap[context] = list
					} else {
						seedMap[context] = append(seedMap[context],newSourceMapper)
					}

				} else if locationRegex.MatchString(line){
					context = "location"
				}
			case "location":
				if numbersRegex.MatchString(line) {
					numbersLoc := getNumbers(line,*numbersRegex)
					newSourceMapper := createSourceMapper(numbersLoc[0],numbersLoc[1],numbersLoc[2])
					_,ok := seedMap[context]
					if (ok == false) {
						var list []sourceMapper
						list = append(list,newSourceMapper)
						seedMap[context] = list
					} else {
						seedMap[context] = append(seedMap[context],newSourceMapper)
					}

				}
		}
	}
	lowest := getNextValue("soil",seedMap,seeds[0])
	
	i := 0
	for {
		if (i < len(seeds) && i+1 < len(seeds)) {
			for x := 0; x < seeds[i+1]; x++{
				if getNextValue("soil",seedMap,seeds[i] + x) < lowest {
					lowest = getNextValue("soil",seedMap,seeds[i] + x)
				}
			}
			i += 2
		} else {
			break
		}
		
	}
	fmt.Println(lowest)	
}

type sourceMapper struct {
	sourceBegin int
	sourceEnd int
	destinyBegin int
}