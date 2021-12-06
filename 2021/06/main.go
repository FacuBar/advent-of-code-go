package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("./data")
	if err != nil {
		panic("couldnt open file")
	}

	scanner := bufio.NewScanner(f)

	scanner.Split(ScanCSV)

	lanternfish := make(map[int]int)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		lanternfish[n] += 1
	}

	for d := 0; d < 256; d++ {
		for f := 0; f <= 9; f++ {
			if f == 0 {
				lanternfish[9] += lanternfish[0]

				lanternfish[7] += lanternfish[0]
				lanternfish[0] -= lanternfish[0]
				continue
			}
			lanternfish[f-1] += lanternfish[f]
			lanternfish[f] -= lanternfish[f]
		}
	}

	total := 0
	for _, v := range lanternfish {
		total += v
	}
	fmt.Println(total)
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
