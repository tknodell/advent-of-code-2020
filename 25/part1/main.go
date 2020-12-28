package main

import (
	"fmt"
	"log"
)

func transformSubjectNumber(subjectNumber, loopSize int) int {
	value := 1
	for i := 0; i < loopSize; i++ {
		value = value * subjectNumber
		value = value % 20201227
	}
	return value
}

func findLoopSize(publicKey int) (int, int) {
	maxSubject := 1000
	maxLoop := 1000
	for subjectNumber := 0; subjectNumber < maxSubject; subjectNumber++ {
		for loopSize := 0; loopSize < maxLoop; loopSize++ {
			if transformSubjectNumber(subjectNumber, loopSize) == publicKey {
				return subjectNumber, loopSize
			}
		}
	}
	log.Fatal("Couldnt find loop size with max of ", maxSubject, maxLoop)
	return -1, -1
}

func main() {

	//  Test inputs
	// cardPublicKey := 5764801
	// doorPublicKey := 17807724

	//  My input
	cardPublicKey := 19774466
	doorPublicKey := 7290641

	fmt.Println("Card public key", cardPublicKey)
	cardSubject, cardLoop := findLoopSize(cardPublicKey)
	fmt.Println("FOUND")
	fmt.Printf("subject number: %v, loop size: %v", cardSubject, cardLoop)

	fmt.Println()
	fmt.Println()

	fmt.Println("Door public key", doorPublicKey)
	doorSubject, doorLoop := findLoopSize(doorPublicKey)
	fmt.Println("FOUND")
	fmt.Printf("subject number: %v, loop size: %v", doorSubject, doorLoop)

	fmt.Println()
	fmt.Println()

	fmt.Println("Encryption key", cardPublicKey, doorLoop)
	fmt.Println(transformSubjectNumber(cardPublicKey, doorLoop))

	fmt.Println("Encryption key", doorPublicKey, cardLoop)
	fmt.Println(transformSubjectNumber(doorPublicKey, cardLoop))

}
