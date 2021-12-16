package main

import (
	"advent-of-code/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// https://adventofcode.com/2021/day/3
func main() {

	pwd, err := os.Getwd()
	helpers.Check(err)
	input := pwd + "\\input"
	file := helpers.InputFileAsSlice(input)

	part1(file)
	part2(file)
}

type commonBit struct {
	mostCommonBit   int
	leastCommonBit  int
	mostCommonRows  []string
	leastCommonRows []string
}

func common(input []string, index int) (output commonBit) {

	var count0, count1 float64
	for i := range input {
		row := strings.Split(input[i], "")
		switch row[index] {
		case "1":
			count1++
		case "0":
			count0++
		}
	}

	switch half := float64(len(input)) / 2; {
	case count1 > half && count0 < half:
		output.mostCommonBit = 1
		output.leastCommonBit = 0
	case count1 == half && count0 == half:
		output.mostCommonBit = 1
		output.leastCommonBit = 0
	case count1 < half && count0 > half:
		output.mostCommonBit = 0
		output.leastCommonBit = 1
	}

	for i := range input {
		row := strings.Split(input[i], "")
		switch row[index] {
		case strconv.Itoa(output.mostCommonBit):
			output.mostCommonRows = append(output.mostCommonRows, input[i])
		case strconv.Itoa(output.leastCommonBit):
			output.leastCommonRows = append(output.leastCommonRows, input[i])
		}
	}
	return output
}

func binArrayToInt(input []int) (output int64) {
	var binary string

	for i := range input {
		binary = binary + strconv.Itoa(input[i])
	}
	output, _ = strconv.ParseInt(binary, 2, 64)
	return output
}

func part1(file []string) {
	var gammaArray, epsilonArray []int

	inputLength := len(strings.Split(file[0], ""))

	for i := 0; i < inputLength; i++ {
		common := common(file, i)
		gammaArray = append(gammaArray, common.mostCommonBit)
		epsilonArray = append(epsilonArray, common.leastCommonBit)
	}
	gammaRate := binArrayToInt(gammaArray)
	fmt.Printf("The Gamma Rate is %d\n", gammaRate)

	epsilonRate := binArrayToInt(epsilonArray)
	fmt.Printf("The Epsilon Rate is %d\n", epsilonRate)
	answer := gammaRate * epsilonRate
	fmt.Printf("The power consumption of the submarine is: %d\n", answer)
}

func part2(file []string) {

	inputLength := len(strings.Split(file[0], ""))
	var oxygenGeneratorArr, cO2ScrubberArr []string
	var listMostCommon, listLeastCommon []string
	listMostCommon = file
	listLeastCommon = file

	for i := 0; i < inputLength; i++ {
		listMostCommon = (common(listMostCommon, i).mostCommonRows)
		if len(listMostCommon) == 1 {
			oxygenGeneratorArr = listMostCommon
			break
		}
	}

	for i := 0; i < inputLength; i++ {
		listLeastCommon = (common(listLeastCommon, i).leastCommonRows)
		if len(listLeastCommon) == 1 {
			cO2ScrubberArr = listLeastCommon
			break
		}
	}
	oxygenGeneratorRating, _ := strconv.ParseInt(oxygenGeneratorArr[0], 2, 64)
	fmt.Printf("The oxygen generator rating is %d\n", oxygenGeneratorRating)
	cO2ScrubberRating, _ := strconv.ParseInt(cO2ScrubberArr[0], 2, 64)
	fmt.Printf("The CO2 scrubber rating is %d\n", cO2ScrubberRating)

	lifeSupportRating := oxygenGeneratorRating * cO2ScrubberRating
	fmt.Printf("The Life Support Rating is %d\n", lifeSupportRating)
}
