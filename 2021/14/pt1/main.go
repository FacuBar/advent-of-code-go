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
	var polymerTemplate []byte
	pairInsertionRules := make([][]string, 0)
	for scanner.Scan() {
		if scanner.Text() == "" {
			isPolymerTemplate = false
			continue
		}

		if isPolymerTemplate {
			polymerTemplate = []byte(scanner.Text())
			continue
		}

		pairInsertionRules = append(pairInsertionRules, strings.Split(scanner.Text(), " -> "))
	}

	var segment []byte
	var sb strings.Builder
	counter := make(map[byte]int, 0)
	for step := 0; step < 10; step++ {
		sb = strings.Builder{}
		for i := 0; i < len(polymerTemplate)-1; i++ {
			segment = polymerTemplate[i : i+2]
			for j, rule := range pairInsertionRules {
				if i == 0 {
					if string(segment) == rule[0] {
						sb.WriteByte(segment[0])
					} else {
						if j == len(pairInsertionRules) {
							sb.WriteByte(segment[0])
						}
					}
				}
				if string(segment) == rule[0] {
					sb.WriteString(rule[1])
					sb.WriteByte(segment[1])
				} else {
					if j == len(pairInsertionRules) {
						sb.WriteByte(segment[1])
					}
				}

			}
		}
		polymerTemplate = []byte(sb.String())
	}

	fmt.Printf("%#v", counter)
	for _, b := range polymerTemplate {
		counter[byte(b)]++
	}

	var max, min, i int
	for k, v := range counter {
		fmt.Println(string(k))
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
