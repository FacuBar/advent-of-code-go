package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("../data")
	if err != nil {
		panic("couldnt open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	nums := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}

	previous := nums[0] + nums[1] + nums[2]
	actual, count := 0, 0

	for i := 0; i < len(nums)-3; i++ {
		actual = nums[i+1] + nums[i+2] + nums[i+3]
		if actual > previous {
			count++
		}
		previous = actual
	}

	fmt.Println(count)
}
