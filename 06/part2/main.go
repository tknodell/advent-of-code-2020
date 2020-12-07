package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	var numInGroup int

	letters := popAlphabet()
	fmt.Println(letters)

	lines, _ := readLines("input.txt")
	for _, v := range lines {

		if v == "" {
			fmt.Println("NEW GROUP, previous had", numInGroup)
			numInGroup = 0
		}

		numInGroup++

		// loop thru letters
		letterCount := popAlphabet()

		for _, l := range v {
			letterString := string(l)
			letterCount[letterString] += 1
			fmt.Println(letterString)
		}
		fmt.Println(letterCount)
		fmt.Println("NEW LINE", v)
	}

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
