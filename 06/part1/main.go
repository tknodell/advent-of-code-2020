package main

import (
	"bufio"
	"fmt"
	"os"
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
	res, _ := readLines("input.txt")
	var group string
	for _, v := range res {
		if v != "" {
			fmt.Println("adding to group", v)
			group += v
		} else {
			fmt.Println("new group; result is", group)
			group = ""
		}
		fmt.Println(v)
	}
}
