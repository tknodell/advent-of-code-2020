package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
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

	for {
		lines, _ := readLines("test.txt")
		loopAround, res := calc(lines)

		if !loopAround {
			fmt.Println()
			fmt.Println("Part 2 answer is:", res)
			os.Exit(0)
		}
	}

}

func calc(lines []string) (bool, int) {

	var indexSeen []int
	var index, accumulator int
	var length int = len(lines)
	var alreadySwapped bool

	for {
		indexSeen = append(indexSeen, index)
		instruction := strings.Split(lines[index], " ")

		num, _ := strconv.Atoi(instruction[1])

		if instruction[0] == "jmp" && !alreadySwapped {
			random := flipCoin()
			if random == 1 {
				fmt.Println("swapped jmp to nop!")
				instruction[0] = "nop"
				alreadySwapped = true
			}
		} else if instruction[0] == "nop" && !alreadySwapped {
			random := flipCoin()
			if random == 1 {
				fmt.Println("swapped nop to jmp!")
				instruction[0] = "jmp"
				alreadySwapped = true
			}
		}

		switch instruction[0] {
		case "acc":
			fmt.Println("Adding", num)
			accumulator += num
		case "jmp":
			fmt.Println("Jumping", num)
			index += num - 1
		case "nop":
			fmt.Println("Skipping")
		}

		index++

		if index == length {
			return false, accumulator
		}

		if contains(indexSeen, index) {
			fmt.Println("Looped back around")
			return true, accumulator
		}
	}
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func flipCoin() int {
	rand.Seed(time.Now().UnixNano())

	if flipint := rand.Intn(2); flipint == 0 {
		return 0
	}
	return 1
}

// 622 too low
