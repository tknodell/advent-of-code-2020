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
	// targetNum := 127
	// lines, _ := readLines("test.txt")

	targetNum := 26796446
	lines, _ := readLines("input.txt")

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines)-i; j++ {
			sumResult := sumSlice(lines, i, j)
			if sumResult == targetNum {
				fmt.Println(lines[i : i+j])
				fmt.Println("FOUND IT", targetNum)
				fmt.Println("ANSWER")
				fmt.Println(smallestPlusLargestNum(lines[i : i+j]))
				os.Exit(0)
			}
		}
	}
}

func sumSlice(n []int, start, length int) int {
	var sum int
	bar := n[start : start+length]
	for _, v := range bar {
		sum += v
	}
	return sum
}

func smallestPlusLargestNum(n []int) int {
	var smallest, largest int
	smallest = n[0] // set to first value
	for _, v := range n {
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}

	// fmt.Println(smallest, largest)
	return smallest + largest

}
