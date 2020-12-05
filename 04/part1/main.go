package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var requiredFields = []string{
	"byr",
	"iyr",
	"eyr",
	"hgt",
	"hcl",
	"ecl",
	"pid",
	// "cid" ,
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

func main() {

	var validPassport int
	lines, _ := readLines("input.txt")
	var Passport string

	for _, v := range lines {

		if v == "" {
			// hit blank line

			if hasValidFields(Passport, requiredFields) {
				fmt.Println("passport is valid")
				fmt.Println(Passport)
				validPassport++
			} else {
				fmt.Println("invalid passport\n")
				fmt.Println(Passport)
			}
			fmt.Println("\n")
			Passport = ""
		} else {
			Passport += v
			// fmt.Println("adding", v)
		}
	}

	fmt.Println(validPassport+1, "valid passports")
}

func hasValidFields(p string, fields []string) bool {
	for _, v := range fields {
		if !strings.Contains(p, v) {
			return false
		}
	}
	return true
}

// 234 too low
// 235
