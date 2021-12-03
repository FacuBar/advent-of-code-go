package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day-1/data")
	if err != nil {
		panic("couldnt open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0

	scanner.Scan()
	previous, _ := strconv.Atoi(scanner.Text())

	for scanner.Scan() {
		x, _ := strconv.Atoi(scanner.Text())
		if x > previous {
			count++
		}
		previous = x
	}

	fmt.Println(count)
}
