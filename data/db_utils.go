package data

import (
	"math/rand"
	"time"
)

const alphaLower string = "abcdefghijklmnopqrstuvwxyz"
const alphaUpper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var randomSeed *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomName(length int) string {
	name := make([]byte, length)

	for i := range name {
		if i == 0 {
			name[i] = alphaUpper[randomSeed.Intn(len(alphaUpper))]
		} else {
			name[i] = alphaLower[randomSeed.Intn(len(alphaLower))]
		}
	}

	return string(name)
}

func GenerateRandomAge(max int, min int) int {
	return randomSeed.Intn(max-min+1) + min
}
