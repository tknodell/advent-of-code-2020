package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		findSeat(v)
	}

	fmt.Println()
	// findSeat("FBFBBFFRLR")
	// findSeat("BFFFBBFRRR")
	// findSeat("FFFBBBFRRR")
	// findSeat("BBFFBBFRLL")

	// Part2
	seatIds, _ := readLines("results2.csv")
	for i, v := range seatIds {
		seatID, _ := strconv.Atoi(v)
		if (i + 75) != seatID {
			fmt.Println("our seat ID is", seatID-1)
			break
		}
	}

}

func findSeat(s string) {
	row := calcRow(s)
	col := calcColumns(s)
	id := seatID(row, col)
	fmt.Println(s, row, col, id)
}

func seatID(r, c int) int {
	return r*8 + c
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

	return rows[0]
}

func calcColumns(s string) int {
	colsOnly := s[len(s)-3:]
	cols := popCols()

	for _, v := range colsOnly {
		letter := string(v)
		switch letter {
		case "R":
			cols = cols[len(cols)/2:]
		case "L":
			cols = cols[0 : len(cols)/2]
		}
	}

	return cols[0]
}

func popCols() []int {
	var r []int
	for i := 0; i < totalCols; i++ {
		r = append(r, i)
	}
	return r
}
