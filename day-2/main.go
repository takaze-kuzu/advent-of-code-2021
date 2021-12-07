package main

import (
	"advent-of-code/helpers"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	pwd, err := os.Getwd()
	helpers.Check(err)
	input := pwd + "\\input"

	file := helpers.InputFileAsSlice(input)
	type vector struct {
		direction string
		distance  int
	}
	var chart []vector
	for i := 0; i < len(file); i++ {
		split := strings.Fields(file[i])
		direction := split[0]
		distance, err := strconv.Atoi(split[1])
		helpers.Check(err)

		a := vector{direction, distance}
		chart = append(chart, a)
	}

	type location struct {
		horizontal int
		depth      int
		aim        int
	}

	var submarineLocation1 location
	var submarineLocation2 location

	for i := 0; i < len(chart); i++ {
		// use case
		switch chart[i].direction {
		case "forward":
			submarineLocation1.horizontal = submarineLocation1.horizontal + chart[i].distance

			submarineLocation2.horizontal = submarineLocation2.horizontal + chart[i].distance
			submarineLocation2.depth = submarineLocation2.depth + (submarineLocation2.aim * chart[i].distance)
		case "down":
			submarineLocation1.depth = submarineLocation1.depth + chart[i].distance

			submarineLocation2.aim = submarineLocation2.aim + chart[i].distance
		case "up":
			submarineLocation1.depth = submarineLocation1.depth - chart[i].distance

			submarineLocation2.aim = submarineLocation2.aim - chart[i].distance
		}
	}
	fmt.Println(submarineLocation1)
	answer1 := submarineLocation1.horizontal * submarineLocation1.depth
	fmt.Println(answer1)

	fmt.Println(submarineLocation2)
	answer2 := submarineLocation2.horizontal * submarineLocation2.depth
	fmt.Println(answer2)
}
