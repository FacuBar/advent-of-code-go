package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coordinates struct {
	Depth      int
	Horizontal int
	Aim        int
}

var (
	position coordinates
)

func main() {
	f, err := os.Open("../data.txt")
	if err != nil {
		panic("couldnt open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	var input string
	for scanner.Scan() {
		input = scanner.Text()
		move(input)
	}

	fmt.Println(position.Depth * position.Horizontal)
}

func move(input string) {
	movement := strings.Split(input, " ")
	value, _ := strconv.Atoi(movement[1])
	switch movement[0] {
	case "forward":
		position.Horizontal += value
		position.Depth += value * position.Aim
		return
	case "up":
		position.Aim -= value
		return
	case "down":
		position.Aim += value
	}

}
