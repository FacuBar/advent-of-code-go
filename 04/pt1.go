package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type bingoNum struct {
	Num    int8
	Marked bool
}

type bingoTable []*bingoNum

func main() {
	f, err := os.Open("./data")
	if err != nil {
		log.Fatalf("couldtn open file, error: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(SplitAt("\n\n"))

	var text string

	scanner.Scan()
	text = scanner.Text()

	draws := make([]int8, 0)
	for _, draw := range strings.Split(text, ",") {
		n, _ := strconv.Atoi(draw)
		draws = append(draws, int8(n))
	}

	bingoTables := make([]bingoTable, 0)
	for scanner.Scan() {
		text = scanner.Text()
		bingoTables = append(bingoTables, getTable(text))
	}
	// Shortcut to resolve a problem induced when deleting tables
	bingoTables = append(bingoTables, bingoTable{})

	for _, draw := range draws {
		for t := 0; t < len(bingoTables); t++ {
			for i, bingoNum := range bingoTables[t] {
				if bingoNum.Num == draw {
					bingoNum.Marked = true

					if isWin(i, bingoTables[t]) {
						fmt.Printf("%v\n", int(draw)*sumUnmarked(bingoTables[t]))
						bingoTables = remove(bingoTables, t)
					}
				}
			}
		}
	}
}

func remove(s []bingoTable, i int) []bingoTable {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func sumUnmarked(table bingoTable) int {
	total := 0
	for _, bingoNum := range table {
		if !bingoNum.Marked {
			total += int(bingoNum.Num)
		}
	}
	return total
}

func isWin(draw int, table bingoTable) bool {
	flag := true
	column := draw % 5
	row := draw - (column)

	for i := column; i < 25; i += 5 {
		if !table[i].Marked {
			flag = false
			break
		}
		flag = true
	}

	if flag {
		return flag
	}

	for i := row; i < row+5; i++ {
		if !table[i].Marked {
			flag = false
			break
		}
		flag = true
	}

	return flag
}

func getTable(text string) bingoTable {
	arrnums := make(bingoTable, 0)

	for _, num := range strings.Split(strings.ReplaceAll(text, "\n", " "), " ") {
		if num == "" {
			continue
		}

		n, _ := strconv.Atoi(num)

		arrnums = append(arrnums, &bingoNum{Num: int8(n)})
	}

	return arrnums
}

func SplitAt(substring string) func(data []byte, atEOF bool) (advance int, token []byte, err error) {
	searchBytes := []byte(substring)
	searchLen := len(searchBytes)
	return func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		dataLen := len(data)

		// Return nothing if at end of file and no data passed
		if atEOF && dataLen == 0 {
			return 0, nil, nil
		}

		// Find next separator and return token
		if i := bytes.Index(data, searchBytes); i >= 0 {
			return i + searchLen, data[0:i], nil
		}

		// If we're at EOF, we have a final, non-terminated line. Return it.
		if atEOF {
			return dataLen, data, nil
		}

		// Request more data.
		return 0, nil, nil
	}
}
