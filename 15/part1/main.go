package main

import (
	"fmt"
)

func main() {
	// input := []int{0, 3, 6} // test input
	input := []int{1, 0, 16, 5, 17, 4}

	// fmt.Println("Answer after 2020 runs", Play(input, 2020))
	fmt.Println("Answer after 30000000 runs", Play(input, 30000000))
}

func Play(starting []int, maxTurns int) int {
	lastSpoken := make(map[int]int)

	var previous int
	speak := starting[0]

	for turn := 0; turn < maxTurns; turn++ {
		previous = speak

		if turn < len(starting) {
			speak = starting[turn] // add inital input values
		} else if _, ok := lastSpoken[speak]; !ok {
			speak = 0
		} else {
			speak = turn - lastSpoken[speak] - 1
		}

		lastSpoken[previous] = turn - 1
		// fmt.Println(speak)
	}

	return speak
}
