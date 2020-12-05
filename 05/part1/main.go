package main

import (
	"bufio"
	"fmt"
	"os"
)

const TotalRows = 128

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	values, _ := readLines("input.txt")
	for _, v := range values {
		rowValue := v[0:7]
		calcRow(rowValue)
	}

	calcRow("FBFBBFFRLR")
}

func popRows() []int {
	var r []int
	for i := 0; i < TotalRows; i++ {
		r = append(r, i)
	}
	return r
}

func calcRow(s string) int {
	rows := popRows()

	for _, v := range s {
		letter := string(v)
		fmt.Println(letter)
		switch letter {
		case "F":
			rows = rows[0 : len(rows)/2]
		case "B":
			rows = rows[len(rows)/2:]
		}
	}

	fmt.Println(rows)
	return len(rows)
}
