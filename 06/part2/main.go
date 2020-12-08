package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var numInGroup int
	var letterCount map[string]int = popAlphabet()
	var totalScore int

	lines, _ := readLines("input.txt")
	for _, v := range lines {

		if v == "" {
			fmt.Println(letterCount)
			fmt.Println("NEW GROUP, previous had", numInGroup)
			score := getValidLetters(letterCount, numInGroup)
			fmt.Println("SCORE IS", score)
			totalScore += score
			numInGroup = 0
			letterCount = popAlphabet()
		} else {
			numInGroup++
		}

		// loop thru letters
		for _, l := range v {
			letterString := string(l)
			letterCount[letterString] += 1
		}

		fmt.Println(v)
	}
	fmt.Println("TOTAL SCORE IS ", totalScore)
}

func getValidLetters(m map[string]int, count int) int {
	var sum int
	for _, v := range m {
		if v == count {
			sum++
		}
	}

	return sum
}

func popAlphabet() map[string]int {
	s := make(map[string]int)
	for _, v := range "abcdefghijklmnopqrstuvwxyz" {
		letterString := string(v)
		s[letterString] = 0
	}
	return s
}

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
