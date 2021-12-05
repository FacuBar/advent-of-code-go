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

	printMap(ventMap)
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
	}
}

func sortAsc(n1 int, n2 int) (int, int) {
	if n1 < n2 {
		return n1, n2
	}
	return n2, n1
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
