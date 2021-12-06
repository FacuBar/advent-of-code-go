package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MAPSIZE = 1000

type ventMap struct {
	Map     *[MAPSIZE][MAPSIZE]int
	overlap uint32
}

func main() {
	f, err := os.Open("../data")
	if err != nil {
		panic("couldnt open file")
	}

	scanner := bufio.NewScanner(f)

	ventMap := &ventMap{
		Map:     &[MAPSIZE][MAPSIZE]int{},
		overlap: 0,
	}

	for scanner.Scan() {
		insertLines(scanner.Text(), ventMap)
	}
	fmt.Println(ventMap.overlap)
}

func insertLines(coord string, ventMap *ventMap) {
	coordinates := strings.Split(coord, " -> ")

	p1 := strings.Split(coordinates[0], ",")
	p1x, _ := strconv.Atoi(p1[0])
	p1y, _ := strconv.Atoi(p1[1])

	p2 := strings.Split(coordinates[1], ",")
	p2x, _ := strconv.Atoi(p2[0])
	p2y, _ := strconv.Atoi(p2[1])

	if p1x == p2x {
		n1, n2 := sortAsc(p1y, p2y)

		for i := n1; i <= n2; i++ {
			ventMap.Map[p1x][i] += 1
			if ventMap.Map[p1x][i] == 2 {
				ventMap.overlap += 1
			}
		}

	} else if p1y == p2y {
		n1, n2 := sortAsc(p1x, p2x)

		for i := n1; i <= n2; i++ {
			ventMap.Map[i][p1y] += 1
			if ventMap.Map[i][p1y] == 2 {
				ventMap.overlap += 1
			}
		}
	} else {
		if p1x == p1y && p2x == p2y {
			n1, n2 := sortAsc(p1x, p2x)
			for i := n1; i <= n2; i++ {
				ventMap.Map[i][i] += 1
				if ventMap.Map[i][i] == 2 {
					ventMap.overlap += 1
				}
			}

		} else if diagonal, direction := isDiagonal(p1x, p1y, p2x, p2y); diagonal {
			x1, y1 := sortPointsByXAsc(p1x, p1y, p2x, p2y)
			n1, n2 := sortAsc(p1x, p2x)
			dif := n2 - n1

			if direction > 0 {
				for i := 0; i <= dif; i++ {
					ventMap.Map[x1+i][y1+i] += 1
					if ventMap.Map[x1+i][y1+i] == 2 {
						ventMap.overlap += 1
					}
				}
				return
			}

			for i := 0; i <= dif; i++ {
				ventMap.Map[x1+i][y1-i] += 1
				if ventMap.Map[x1+i][y1-i] == 2 {
					ventMap.overlap += 1
				}
			}
		}
	}
}

// returns the point with the lowest x value
func sortPointsByXAsc(points ...int) (int, int) {
	if points[0] < points[2] {
		return points[0], points[1]
	}
	return points[2], points[3]
}

func sortAsc(n1 int, n2 int) (int, int) {
	if n1 < n2 {
		return n1, n2
	}
	return n2, n1
}

// determines wether two points forms part of a
// segment with a 45 degree inclinaction or not
func isDiagonal(nums ...int) (bool, int) {
	r1 := nums[0] - nums[2]
	r2 := nums[1] - nums[3]

	return r1 == r2 || r1 == -r2, r1 / r2
}

// For debuggin purpouses only
func printMap(ventMap *ventMap) {
	for x := range ventMap.Map {
		for y := range ventMap.Map {
			fmt.Print(" ", ventMap.Map[y][x])
		}
		fmt.Println("")
	}
}
