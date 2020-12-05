package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
				if hasValidValues(Passport) {
					fmt.Println("passport is valid")
					// fmt.Println(Passport)
					validPassport++
				}
			}
			fmt.Println()
			Passport = ""
		} else {
			Passport += v + " "
		}
	}

	fmt.Println(validPassport, "valid passports")
}

func hasValidFields(p string, fields []string) bool {
	for _, v := range fields {
		if !strings.Contains(p, v) {
			return false
		}
	}
	return true
}

func hasValidValues(p string) bool {
	fmt.Println("Eval", p)
	passportSlice := strings.Split(p, " ")

	for _, pass := range passportSlice {
		if pass != "" {
			fieldSlice := strings.Split(pass, ":")
			fmt.Println(fieldSlice[0], fieldSlice[1])
			if !validValue(fieldSlice) {
				return false
			}
		}
	}
	return true
}

func validValue(s []string) bool {
	switch s[0] {
	case "byr":
		val, _ := strconv.Atoi(s[1])
		if !(len(s[1]) == 4 && val >= 1920 && val <= 2002) {
			return false
		}
	case "iyr":
		val, _ := strconv.Atoi(s[1])
		if !(len(s[1]) == 4 && val >= 2010 && val <= 2020) {
			return false
		}
	case "eyr":
		val, _ := strconv.Atoi(s[1])
		if !(len(s[1]) == 4 && val >= 2020 && val <= 2030) {
			return false
		}
	case "hgt":
		val := s[1]
		fmt.Println(val)

		suffix := val[len(val)-2:]
		value, err := strconv.Atoi(strings.Split(val, suffix)[0])
		if err != nil {
			fmt.Println(err)
		}

		switch suffix {
		case "cm":
			if !(value >= 150 && value <= 193) {
				return false
			}
		case "in":
			if !(value >= 59 && value <= 76) {
				return false
			}
		default:
			return false
		}
	case "hcl":
		val := s[1]
		match, _ := regexp.MatchString("#([0-9]|[a-f]){6}", val)
		return match
	case "ecl":
		val := s[1]
		match, _ := regexp.MatchString("amb|blu|brn|gry|grn|hzl|oth", val)
		return match
	case "pid":
		val := s[1]
		match, _ := regexp.MatchString("[0-9]{9}", val)
		return match
	}

	return true
}

// 196 too high
// 195 too high
