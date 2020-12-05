package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		fmt.Println(calcRow(rowValue))
	}
}

func calcRow(s string) int {
	return -1
}
