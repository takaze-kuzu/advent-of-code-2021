package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"advent-of-code/helpers"
)

func main() {
	part1("file")
	part2("file")
}

func part1(input string) {
	f, err := os.Open(input)
	helpers.Check(err)
	defer f.Close()

	var p, c, count int = 0, 0, -1
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		c, err = strconv.Atoi(scanner.Text())
		helpers.Check(err)
		if c > p {
			count++
		}
		p = c
	}
	fmt.Println(count)
}

func part2(input string) {
	f, err := os.Open(input)
	helpers.Check(err)
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var arr []int
	for scanner.Scan() {
		element, err := strconv.Atoi(scanner.Text())
		helpers.Check(err)
		arr = append(arr, element)
	}

	var c, s, previous, count = 0, 3, 0, -1
	for c+s <= len(arr) {
		current := 0
		for i := 0; i < s; i++ {
			current = current + arr[c+i]
		}
		if current > previous {
			count++
		}
		c++
		previous = current
	}
	fmt.Println(count)
}
