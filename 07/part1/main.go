package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
	bagFormat := regexp.MustCompile(`(\d+) (\w+).*(bags?)`)

	lines, _ := readLines("test.txt")

	outerBags := make(map[string]int)
	// Get outer bags
	for _, v := range lines {
		bags := strings.Split(v, "bags contain")
		outerBags[bags[0]] = 0
	}

	// Get inner bags
	for _, v := range lines {
		bags := strings.Split(v, "bags contain")
		if bags[1] != " no other bags." {
			inner := bagFormat.FindAllString(bags[1], -1)
			fmt.Println(bags[0])
			// fmt.Println(inner[0])

			// Split on comma
			bag := strings.Split(inner[0], ",")

			for _, b := range bag {
				// Get our number to add
				fmt.Println(strings.TrimSpace(b))

			}
			fmt.Println()
		}
	}

	// for i, v := range outerBags {
	// 	fmt.Println(i, v)
	// }

}
