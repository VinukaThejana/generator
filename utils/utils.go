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

	randomNumRange := [2]int{0, (len(chars) - 1)}

	iterations := getRandomInt(randomNumRange[0], randomNumRange[1])
	replacePoints := []int{}
	for i := 0; i < iterations; i++ {
		replacePoints = append(replacePoints, getRandomInt(randomNumRange[0], randomNumRange[1]))
	}

	for k := 0; k < iterations; k++ {
		chars[k] = getRandomSpecialCharacters()
	}

	return string(chars)
}
