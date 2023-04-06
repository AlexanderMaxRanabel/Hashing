package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

func block_info() {
	rand.Seed(time.Now().UnixNano())
	// Define the characters that can be used in the random string
	const chars = "0123456789"
	var blockID string
	for i := 0; i < 10; i++ {
		blockID += string(chars[rand.Intn(len(chars))])
	}

	fmt.Println("Block Has been generated: ID", blockID)
}
func main() {
	// Initialize random number generator with a seed based on the current time
	rand.Seed(time.Now().UnixNano())

	// Define the characters that can be used in the random string
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	// Generate a random string of length 10
	var result string
	for i := 0; i < 10; i++ {
		result += string(chars[rand.Intn(len(chars))])
	}

	fmt.Print("Enter a data to hash: ")
	var data string
	fmt.Scan(&data)

	result_hash := sha256.Sum256([]byte(result))
	hash_1 := sha256.Sum256([]byte(data))
	hash_2 := sha256.Sum256(hash_1[:])
	hash_3 := sha256.Sum256(hash_2[:])

	if bytes.Equal(result_hash[:], hash_3[:]) {
		fmt.Println("Hashes mach. Block cannot be generated")
	} else {
		block_info()
		fmt.Println("Data: ", hash_3)
		fmt.Println("Target: ", result_hash)
		for {
			hehe := 0
			time.Sleep(time.Duration(hehe) * time.Second) // Pause execution for 5 seconds
			hehe = hehe + 1
		}
	}
}
