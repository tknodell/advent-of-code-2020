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
	preambleLength := 5
	preambleStart := 0

	lines, _ := readLines("test.txt")

	preamble := getPreamble(lines, preambleStart, preambleLength)
	fmt.Println(validNextNum(preamble, 40))
}

func getPreamble(n []int, start, length int) []int {
	return n[start : start+length]
}

func validNextNum(n []int, num int) bool {
	for _, i := range n {
		for _, j := range n {
			if i+j == num && i != j {
				return true
			}
		}
	}

	return false
}
