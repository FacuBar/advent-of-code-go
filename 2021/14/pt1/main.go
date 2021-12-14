package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("../data")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	isPolymerTemplate := true
	var polymerTemplate string
	pairInsertionRules := make([][]string, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			isPolymerTemplate = false
			continue
		}

		if isPolymerTemplate {
			polymerTemplate = scanner.Text()
			continue
		}

		pairInsertionRules = append(pairInsertionRules, strings.Split(scanner.Text(), " -> "))
	}

	pairCounter := make(map[string]int, 0)
	for i := 0; i < len(polymerTemplate)-1; i++ {
		pairCounter[string(polymerTemplate[i:i+2])]++
	}

	letterCounter := make(map[byte]int, 0)

	for step := 0; step < 40; step++ {
		tempCounter := make(map[string]int, 0)
		for k, v := range pairCounter {
			for _, rule := range pairInsertionRules {
				if k == rule[0] {
					tempCounter[string(k[0])+rule[1]] += v
					tempCounter[rule[1]+string(k[1])] += v
					break
				}
			}
		}
		pairCounter = tempCounter
	}

	// given that every pair overlaps the first and last letter
	// only makes sense to count either the first or last letter
	// of each pair; and add the first o last letter of the original
	// template given that the extrems never change
	for k, v := range pairCounter {
		letterCounter[k[0]] += v
	}
	letterCounter[polymerTemplate[len(polymerTemplate)-1]]++

	var max, min, i int

	for _, v := range letterCounter {
		if i == 0 {
			max, min = v, v
			i++
			continue
		}

		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	fmt.Println(max - min)
}
