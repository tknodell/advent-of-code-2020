package main

import (
	"fmt"
	"math/big"
)

func main() {
	//  Test inputs
	// cardPublicKey := 5764801
	// doorPublicKey := 17807724

	//  My input
	cardPublicKey := 19774466
	doorPublicKey := 7290641

	value := 1
	loop_size := 0

	for value != cardPublicKey {
		value = (value * 7) % 20201227
		loop_size++
	}

	fmt.Println("Found loop size:", loop_size)

	var encryptionKey = new(big.Int)
	encryptionKey.Exp(
		big.NewInt(int64(doorPublicKey)),
		big.NewInt(int64(loop_size)),
		big.NewInt(20201227))

	fmt.Println("Encryption key:", encryptionKey)
}
