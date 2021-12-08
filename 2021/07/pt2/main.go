package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("../data")

	scanner := bufio.NewScanner(f)

	scanner.Split(ScanCSV)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	min, max := n, n

	crabs := make([]int, 0)
	for ok := true; ok; ok = scanner.Scan() {
		n, _ = strconv.Atoi(scanner.Text())

		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
		crabs = append(crabs, n)
	}

	var val, sum int
	for position := min; position <= max; position++ {
		sum = 0
		for _, crab := range crabs {
			sum += (abs(crab-position) * (abs(crab-position) + 1)) / 2
		}
		if position == min {
			val = sum
		} else {
			if sum < val {
				val = sum
			}
		}
	}

	fmt.Println(val)
}

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func ScanCSV(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find next separator and return token
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}

	// Request more data.
	return 0, nil, nil
}
