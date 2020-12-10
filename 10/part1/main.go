package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	lines, _ := readLines("input.txt")

	lines = append(lines, 0)
	sort.Ints(lines)

	var differences []int

	for i := 0; i < len(lines)-1; i++ {
		jolt := lines[i]
		difference := lines[i+1] - jolt
		differences = append(differences, difference)
		// fmt.Println(difference)
	}
	// add adapter
	differences = append(differences, 3)

	fmt.Println(differences)

	fmt.Println(countOccurances(differences, 1))
	fmt.Println(countOccurances(differences, 3))

}

func countOccurances(n []int, t int) int {
	var count int
	for _, v := range n {
		if v == t {
			count++
		}
	}
	return count
}
