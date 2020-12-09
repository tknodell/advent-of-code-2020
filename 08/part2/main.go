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
	var count int

	for {
		lines, _ := readLines("input.txt")
		loopAround, res := calc(lines, count)

		if !loopAround {
			fmt.Println()
			fmt.Println("Part 2 answer is:", res)
			os.Exit(0)
		}
		count++
	}

}

func calc(lines []string, count int) (bool, int) {

	var indexSeen []int
	var index, accumulator int
	var length int = len(lines)

	for {
		indexSeen = append(indexSeen, index)
		instruction := strings.Split(lines[index], " ")

		num, _ := strconv.Atoi(instruction[1])

		if instruction[0] == "jmp" && index == count {
			fmt.Println("swapped jmp to nop! at", count)
			fmt.Println(index, count)
			instruction[0] = "nop"
		}
		// if dont get answer with above re-run with this
		// else if instruction[0] == "nop" {
		// 	fmt.Println("swapped nop to jmp!")
		// 	instruction[0] = "jmp"
		// }

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

// 622 too low
// 4731 too high
