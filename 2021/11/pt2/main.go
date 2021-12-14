package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	total = 0
)

func main() {
	f, _ := os.Open("../data")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	grid := make([][]int, 0)
	for scanner.Scan() {
		row := make([]int, 0)
		for _, chr := range scanner.Text() {
			n, _ := strconv.Atoi(string(chr))
			row = append(row, n)
		}
		grid = append(grid, row)
	}

	for i := 0; i > -1; i++ {
		for row := 0; row < len(grid); row++ {
			for column := 0; column < len(grid); column++ {
				grid[row][column]++

				if grid[row][column] == 10 {
					grid[row][column]++
					total += 1
					flash(grid, row, column)
				}
			}
		}
		// resets octupuses; if all have done so, finishes program
		if allFlash(grid) {
			fmt.Println(i + 1)
			return
		}
	}

	fmt.Println(total)
}

func flash(grid [][]int, row, column int) {
	for _, i := range []int{row - 1, row, row + 1} {
		if i < 0 || i >= len(grid) {
			continue
		}
		for _, j := range []int{column - 1, column, column + 1} {
			if j < 0 || j >= len(grid) {
				continue
			}
			grid[i][j]++
			if grid[i][j] == 10 {
				total += 1
				flash(grid, i, j)
			}
		}
	}
}

func allFlash(grid [][]int) bool {
	flag := true
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[i][j] > 9 {
				grid[i][j] = 0
			} else {
				flag = false
			}
		}
	}
	return flag
}

// debug purpouses only
func printGrid(grid [][]int) {
	for _, row := range grid {
		fmt.Printf("%#v\n", row)
	}
	fmt.Println("")
}
