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
	lines, _ := readLines("test.txt")

	lines = append(lines, 0)

	// add adapter
	lines = append(lines, 3)

	sort.Ints(lines)

	origLen := len(lines)
	fmt.Println(lines)
	for i := 0; i < origLen; i++ {
		fmt.Println(lines)
		fmt.Println(validJoltage(lines))
		lines = lines[1:]
	}

	fmt.Println(validJoltage(lines))

}

func validJoltage(sequence []int) bool {
	sort.Ints(sequence)
	// fmt.Println(sequence)

	for j := 0; j < len(sequence)-1; j++ {
		difference := sequence[j+1] - sequence[j]
		fmt.Println(difference)
		if difference > 3 {
			return false
		}
		// fmt.Println(sequence[j])
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
