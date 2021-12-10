package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	openingChar = map[rune]bool{
		'(': true,
		'[': true,
		'{': true,
		'<': true,

		')': false,
		']': false,
		'}': false,
		'>': false,
	}

	pairChar = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
		'<': '>',
	}

	charValue = map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}
)

func main() {
	f, _ := os.Open("../data")

	scanner := bufio.NewScanner(f)

	results := make([]int, 0)
	for scanner.Scan() {
		result := completeLine(scanner.Text())
		if result != 0 {
			results = append(results, result)
		}
	}
	sort.Ints(results)
	fmt.Println(results[len(results)/2])
}

func completeLine(line string) int {
	sequence := make([]rune, 0)
	opening := ' '

	for _, rn := range line {
		if openingChar[rn] {
			sequence = append(sequence, rn)
		} else {
			opening, sequence = pop(sequence)
			if rn != pairChar[opening] {
				return 0
			}
		}
	}

	if len(sequence) == 0 {
		return 0
	}

	total := 0
	for i := len(sequence) - 1; i > -1; i-- {
		total = total*5 + charValue[sequence[i]]
	}

	return total
}

// using slice as if it were a stack
func pop(sequence []rune) (rune, []rune) {
	char := sequence[len(sequence)-1]
	sequence = sequence[:len(sequence)-1]
	return char, sequence
}
