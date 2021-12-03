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
}

var (
	position coordinates
)

func main() {
	f, err := os.Open("02/data.txt")
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
		return
	case "up":
		position.Depth -= value
		return
	case "down":
		position.Depth += value
	}

}
