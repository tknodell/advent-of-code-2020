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

	letter := strings.TrimSpace(strings.Replace(splitUp[1], ":", "", 1))

	password := splitUp[2]

	lowNumInt, _ := strconv.Atoi(lowNum)
	highNumInt, _ := strconv.Atoi(strings.TrimSpace(highNum))

	letterSlice := password[lowNumInt-1 : highNumInt]

	letterAtLowNum := string(letterSlice[0])
	letterAtHighNum := string(letterSlice[len(letterSlice)-1])

	// fmt.Println(password, letterSlice, lowNumInt, letterAtLowNum, highNumInt, letterAtHighNum)

	if (letter == letterAtLowNum || letter == letterAtHighNum) &&
		(letterAtLowNum != letterAtHighNum) {
		fmt.Println(password, "is valid against", splitUp)
		return true
	}

	return false
}
