package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("../data")

	scanner := bufio.NewScanner(f)

	instances := make(map[int]int)
	var t string
	for scanner.Scan() {
		t = scanner.Text()
		extractPatters(t, instances)
	}

	fmt.Printf("%#v\n", instances)
	// count of 1, 4, 7, 8
	total := instances[2] + instances[4] + instances[3] + instances[7]
	fmt.Println(total)
}

func extractPatters(t string, instances map[int]int) {
	output := strings.Split(t, " | ")
	for _, sequence := range strings.Split(output[1], " ") {
		instances[len(sequence)] += 1
	}
}
