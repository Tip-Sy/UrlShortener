package main

import "math/rand"

// Note: these functions seems to generate the same random string/number from one execution to another

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randomString(length int) string {
	buffer := make([]rune, length)
	for i := range buffer {
		buffer[i] = letters[rand.Intn(len(letters))]
	}
	return string(buffer)
}

func randomNumber(length int) int {
	return rand.Intn(length)
}
