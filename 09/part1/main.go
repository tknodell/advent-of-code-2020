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
	preambleLength := 25
	preambleStart := 0

	lines, _ := readLines("input.txt")

	for index, line := range lines {
		// make sure we're not in the preamble
		if index >= preambleLength {
			preamble := getPreamble(lines, preambleStart, preambleLength)
			isValidNum := validNextNum(preamble, line)
			fmt.Println(line, isValidNum, preamble)
			if !isValidNum {
				os.Exit(0)
			}
			preambleStart++
		}
	}
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

// 229 not right
