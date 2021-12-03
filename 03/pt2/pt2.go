package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("03/data")
	if err != nil {
		panic("couldnt open file")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	binaries := make([]string, 0)
	var bits string
	for scanner.Scan() {
		bits = scanner.Text()
		binaries = append(binaries, bits)
	}

	oxigen := make(chan int64)
	notoxigen := make(chan int64)

	go getRating(binaries, true, oxigen)
	go getRating(binaries, false, notoxigen)

	a := <-oxigen
	b := <-notoxigen

	fmt.Println(a * b)
}

func getRating(binaries []string, isOxygen bool, response chan int64) {
	var bits string
	var counter int
	var filterBy byte
	var val int64

	for i := 0; i < len(binaries[0]); i++ {
		counter = 0

		for _, bits = range binaries {
			switch bits[i] {
			case '1':
				counter++
				break
			case '0':
				counter--
			}
		}

		if isOxygen {
			if counter >= 0 {
				filterBy = '1'
			} else {
				filterBy = '0'
			}
		} else {
			if counter >= 0 {
				filterBy = '0'
			} else {
				filterBy = '1'
			}
		}

		binaries = filter(binaries, filterBy, i)

		if len(binaries) == 1 {
			val, _ = strconv.ParseInt(binaries[0], 2, 64)
			response <- val
			return
		}
	}
	val, _ = strconv.ParseInt(binaries[0], 2, 64)
	response <- val
}

func filter(binaries []string, filterBy byte, i int) []string {
	newBinaries := make([]string, 0)
	var bits string

	for _, bits = range binaries {
		if bits[i] == filterBy {
			newBinaries = append(newBinaries, bits)
		}
	}

	return newBinaries
}
