package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	var validPassCount int
	lines, _ := readLines("input.txt")
	for _, v := range lines {
		if isValidPass(v) {
			validPassCount++
		}
	}
	fmt.Println(validPassCount, "valid passwords")
}

func isValidPass(s string) bool {
	splitUp := strings.SplitAfter(s, " ")

	numRange := splitUp[0]
	lowNum := strings.Split(numRange, "-")[0]
	highNum := strings.Split(numRange, "-")[1]

	letter := strings.Replace(splitUp[1], ":", "", 1)
	// fmt.Println(letter)

	password := splitUp[2]
	letterCount := strings.Count(password, strings.TrimSpace(letter))

	lowNumInt, _ := strconv.Atoi(lowNum)
	highNumInt, _ := strconv.Atoi(strings.TrimSpace(highNum))

	// fmt.Println(letter, password, letterCount, lowNumInt, highNumInt)

	if letterCount >= lowNumInt && letterCount <= highNumInt {
		fmt.Println(password, "is valid")
		return true
	}

	return false
}
