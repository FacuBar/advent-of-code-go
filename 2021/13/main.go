package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

func main() {
	f, _ := os.Open("data")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	paper := make(map[coordinates]bool)

	var folds bool
	var x, y int
	var coords, foldInstructs []string
	for scanner.Scan() {
		if scanner.Text() == "" {
			folds = true
			continue
		}

		if folds {
			foldInstructs = append(foldInstructs, scanner.Text())
			continue
		}

		coords = strings.Split(scanner.Text(), ",")
		x, _ = strconv.Atoi(coords[0])
		y, _ = strconv.Atoi(coords[1])

		paper[coordinates{x: x, y: y}] = true
	}

	var fold []string
	var foldAlong string
	var foldValue int
	for i := 0; i < len(foldInstructs); i++ {
		fold = strings.Split(strings.Split(foldInstructs[i], " ")[2], "=")
		foldAlong = fold[0]
		foldValue, _ = strconv.Atoi(fold[1])
		for coord := range paper {
			switch foldAlong {
			case "x":
				if coord.x > foldValue {
					x = foldValue - (coord.x - foldValue)
					paper[coordinates{x: x, y: coord.y}] = true
					delete(paper, coord)
				}
			case "y":
				if coord.y > foldValue {
					y = foldValue - (coord.y - foldValue)
					paper[coordinates{x: coord.x, y: y}] = true
					delete(paper, coord)
				}
			}
		}
		if i == 0 {
			fmt.Println("pt1: ", len(paper))
		}
	}
	fmt.Println("pt2: ----------")
	plotPaper(paper)
}

func plotPaper(paper map[coordinates]bool) {
	var maxX, maxY int

	for coord := range paper {
		if coord.x > maxX {
			maxX = coord.x
		}

		if coord.y > maxY {
			maxY = coord.y
		}
	}

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if paper[coordinates{x: x, y: y}] {
				fmt.Print("x")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println("")
	}
}
