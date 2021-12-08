package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type sequence struct {
	common7 int
	common4 int
	str     string

	value int
}

var (
	uniqueLenByValue = map[int]int{
		1: 2,
		7: 3,
		4: 4,
		8: 7,
	}
)

func main() {
	f, _ := os.Open("../data")

	scanner := bufio.NewScanner(f)
	total := 0
	for scanner.Scan() {
		total += getOutput(scanner.Text())
	}
	fmt.Println(total)
}

func getOutput(text string) int {
	sequences := make(map[int][]*sequence)

	input := strings.Split(text, " | ")
	for _, sqc := range strings.Split(input[0], " ") {
		sequences[len(sqc)] = append(sequences[len(sqc)], &sequence{str: sqc})
	}

	sequences[2][0].value = 1
	sequences[3][0].value = 7
	sequences[4][0].value = 4
	sequences[7][0].value = 8

	for _, len := range []int{5, 6} {
		decodeSequence(sequences[len], sequences)
	}

	output := 0
	var lensqc int

	for i, sqc := range strings.Split(input[1], " ") {
		lensqc = len(sqc)
		if lensqc != 5 && lensqc != 6 {
			output += sequences[lensqc][0].value * int(math.Pow10(3-i))
			continue
		}

		output += getSqcValue(sqc, sequences[lensqc]) * int(math.Pow10(3-i))
	}
	return output
}

func getSqcValue(sqcstr string, sqcs []*sequence) int {
	for _, sqc := range sqcs {
		if commonLetters(sqc.str, sqcstr) == len(sqcstr) {
			return sqc.value
		}
	}
	return 0
}

func decodeSequence(sqcs []*sequence, sequences map[int][]*sequence) {
	for _, sqc := range sqcs {
		sqc.common4 = commonLetters(sqc.str, sequences[uniqueLenByValue[4]][0].str)
		sqc.common7 = commonLetters(sqc.str, sequences[uniqueLenByValue[7]][0].str)

		if len(sqc.str) == 6 {
			if sqc.common4 == 3 {
				if sqc.common7 == 2 {
					sqc.value = 6
					continue
				}
				sqc.value = 0
				continue
			}
			sqc.value = 9
		} else {
			if sqc.common4 == 3 {
				if sqc.common7 == 2 {
					sqc.value = 5
					continue
				}
				sqc.value = 3
				continue
			}
			sqc.value = 2
		}
	}
}

func commonLetters(s1 string, s2 string) int {
	total := 0

	for _, letter1 := range s1 {
		for _, letter2 := range s2 {
			if letter1 == letter2 {
				total += 1
			}
		}
	}
	return total
}
