package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, value)
	}
	return lines, scanner.Err()
}

func main() {
	bar, _ := readLines("input.txt")

	for _, i := range bar {
		for _, j := range bar {
			if i+j == 2020 {
				fmt.Println(i, j)
				fmt.Println(i * j)
				os.Exit(0)
			}
		}
	}
}
