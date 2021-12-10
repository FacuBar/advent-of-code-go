package main

import (
	"bufio"
	"fmt"
	"os"
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
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
)

func main() {
	f, _ := os.Open("../data")

	scanner := bufio.NewScanner(f)

	total := 0
	for scanner.Scan() {
		total += validateLine(scanner.Text())
	}

	fmt.Println(total)
}

func validateLine(line string) int {
	sequence := make([]rune, 0)
	opening := ' '
	for _, rn := range line {
		if openingChar[rn] {
			sequence = append(sequence, rn)
		} else {
			opening, sequence = pop(sequence)
			if rn != pairChar[opening] {
				return charValue[rn]
			}
		}
	}
	return 0
}

// using slice as if it were a stack
func pop(sequence []rune) (rune, []rune) {
	char := sequence[len(sequence)-1]
	sequence = sequence[:len(sequence)-1]
	return char, sequence
}

// debugging purpouses only
func printSequence(sequence []rune) {
	for _, rune := range sequence {
		fmt.Print(string(rune), " ")
	}
	fmt.Println()

}
