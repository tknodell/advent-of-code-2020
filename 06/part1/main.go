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
	var countSum int
	var groupIndex int
	var validGroup bool = true
	for _, v := range res {
		if v != "" {
			if groupIndex == 0 {
				fmt.Println("adding first value", v)
				group += v
				groupIndex++
				fmt.Println(groupIndex, v)
			} else {
				if !stringContains(group, v) {
					fmt.Printf("%s doesnt contain %v skipping \n", group, v)
					validGroup = false
				} else {
					if stringContains(v, group) {
						fmt.Println("adding to group", v)
						group += v
					}
				}
			}
		} else {
			fmt.Println("new group; result is", group)
			if validGroup {
				fmt.Println("THIS IS RIGHT")
				countSum += uniqLetters(group)
			}
			group = ""
			groupIndex = 0
			validGroup = true
		}
		fmt.Println(v)
	}
	fmt.Println("result is ", countSum)
}

func uniqLetters(s string) int {
	var letters []string

	for _, v := range s {
		value := string(v)
		if !contains(letters, value) {
			letters = append(letters, value)
		}
		// fmt.Println(string(v))
	}
	fmt.Println("uniq letters", len(letters))
	return len(letters)
}

func stringContains(s string, e string) bool {
	for _, a := range s {
		if string(a) == e {
			return true
		}
	}
	return false
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// 46 not right
// 25 not right
