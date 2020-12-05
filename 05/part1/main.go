package main

import (
	"bufio"
	"fmt"
	"os"
)

const totalRows = 128
const totalCols = 8

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
		calcRow(v)
	}

	calcRow("FBFBBFFRLR")
	calcColumns("FBFBBFFRLR")
}

func popRows() []int {
	var r []int
	for i := 0; i < totalRows; i++ {
		r = append(r, i)
	}
	return r
}

func calcRow(s string) int {
	rowsOnly := s[0:7]
	rows := popRows()

	for _, v := range rowsOnly {
		letter := string(v)
		switch letter {
		case "F":
			rows = rows[0 : len(rows)/2]
		case "B":
			rows = rows[len(rows)/2:]
		}
	}

	fmt.Println(s, rows[0])
	return rows[0]
}

func calcColumns(s string) int {
	cols := popCols()

	for _, v := range s {
		letter := string(v)
		switch letter {
		case "R":
			cols = cols[0 : len(cols)/2]
		case "L":
			cols = cols[len(cols)/2:]
		}
	}

	fmt.Println(s, cols[0])
	return cols[0]
}

func popCols() []int {
	var r []int
	for i := 0; i < totalCols; i++ {
		r = append(r, i)
	}
	return r
}
