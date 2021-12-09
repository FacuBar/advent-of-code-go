package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("../data")

	scanner := bufio.NewScanner(f)

	table := make([][]int, 0)
	i := 0
	var row string

	for scanner.Scan() {
		row = scanner.Text()
		table = append(table, make([]int, 0))
		for j := 0; j < len(row); j++ {
			n, _ := strconv.Atoi(string(row[j]))
			table[i] = append(table[i], n)
		}
		i++
	}

	var celVal, counter int
	var flag bool
	for i := 0; i < len(table); i++ {
		for j := 0; j < len(table[i]); j++ {
			flag = true
			celVal = table[i][j]

			if j+1 != len(table[i]) {
				if table[i][j+1] <= celVal {
					flag = false
				}
			}
			if j-1 != -1 {
				if table[i][j-1] <= celVal {
					flag = false
				}
			}

			if i+1 != len(table) {
				if table[i+1][j] <= celVal {
					flag = false
				}
			}
			if i-1 != -1 {
				if table[i-1][j] <= celVal {
					flag = false
				}
			}

			if flag {
				counter += 1 + celVal
			}
		}
	}
	fmt.Println(counter)
}
