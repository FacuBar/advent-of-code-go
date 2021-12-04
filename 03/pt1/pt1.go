package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("../data")
	if err != nil {
		panic("couldnt open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Scan()
	first := scanner.Text()

	binLen := len(first)
	bitCounter := make([]int, binLen)

	var bits string
	for ok := true; ok; ok = scanner.Scan() {
		bits = scanner.Text()

		for i := 0; i < len(bits); i++ {
			switch bits[i] {
			case '1':
				bitCounter[i]++
				break
			case '0':
				bitCounter[i]--
			}
		}
	}

	delta := make([]string, binLen)
	gamma := make([]string, binLen)

	for i := 0; i < len(bitCounter); i++ {
		if bitCounter[i] > 0 {
			delta[i] = "1"
			gamma[i] = "0"
		} else {
			delta[i] = "0"
			gamma[i] = "1"
		}
	}

	deltaV, _ := strconv.ParseInt(strings.Join(delta, ""), 2, 64)
	gammaV, _ := strconv.ParseInt(strings.Join(gamma, ""), 2, 64)
	fmt.Println(deltaV * gammaV)
}
