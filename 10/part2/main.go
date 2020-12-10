package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	combo "github.com/mxschmitt/golang-combinations"
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
		// value, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	lines, _ := readLines("test.txt")

	lines = append(lines, "0")
	// sort.Ints(lines)

	fmt.Println(lines)

	combinations := combo.All(lines)
	fmt.Println(len(combinations))

	var validCombo int

	// var differences []int
	for i := 0; i < len(combinations); i++ {
		com := stringsToInts(combinations[i])

		// add adapter
		com = append(com, 3)

		validJoltage(com)
		fmt.Println(validJoltage(com))
	}

	fmt.Println(validCombo)
}

func validJoltage(sequence []int) bool {
	sort.Ints(sequence)
	fmt.Println(sequence)

	for j := 0; j < len(sequence)-1; j++ {
		difference := sequence[j+1] - sequence[j]
		if difference > 3 {
			return false
		}
		fmt.Println(sequence[j])
	}
	return true
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

func stringsToInts(t []string) []int {
	var t2 = []int{}

	for _, i := range t {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		t2 = append(t2, j)
	}
	return t2
	// fmt.Println(t2)
}
