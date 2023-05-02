/*
* Copyright © 2023 Vinuka Kodituwakku <vinuka.t@pm.me>
 */

// Package utils is a package that contains various functions
// that are being used globally
package utils

import (
	"math/rand"
	"time"
)

func getRandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func getRandomLetter() string {
	rand.Seed(time.Now().UnixNano())
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	return string(alphabet[rand.Intn(len(alphabet))])
}

func getRandomSpecialCharacters() rune {
	specialChars := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', ']', '{', '}', '|', '\\', ';', ':', '"', '\'', ',', '<', '.', '>', '/', '?'}
	rand.Seed(time.Now().UnixNano())
	return specialChars[rand.Intn(len(specialChars))]
}

func randomizeWithSymbols(str string) string {
	chars := []rune(str)
	if len(chars) <= 2 {
		return str
	}

	// Random number with the range
	randomNumRange := struct {
		Min int
		Max int
	}{
		Min: 0,
		Max: len(chars),
	}

	iterations := getRandomInt(randomNumRange.Min, randomNumRange.Max-1)
	replacePoints := []int{}
	for i := 0; i < iterations; i++ {
		replacePoints = append(replacePoints, getRandomInt(randomNumRange.Min, randomNumRange.Max-1))
	}

	for _, index := range replacePoints {
		chars[index] = getRandomSpecialCharacters()
	}

	return string(chars)
}
