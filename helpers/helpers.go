package helpers

import (
	"bufio"
	"os"
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func InputFileAsSlice(f string) (o []string) {
	file, err := os.Open(f)
	Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		o = append(o, scanner.Text())
	}
	return o
}
