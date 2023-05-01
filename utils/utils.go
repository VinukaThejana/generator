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

func getRandom(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func randomizeWithSymbols(str string) string {
	specialChars := []rune{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '=', '+', '[', ']', '{', '}', '|', '\\', ';', ':', '"', '\'', ',', '<', '.', '>', '/', '?'}

	chars := []rune(str)
	if len(chars) <= 2 {
		return str
	}

	randomNumRange := [2]int{0, (len(chars) - 1)}

	iterations := getRandom(randomNumRange[0], randomNumRange[1])
	replacePoints := []int{}
	for i := 0; i < iterations; i++ {
		replacePoints = append(replacePoints, getRandom(randomNumRange[0], randomNumRange[1]))
	}

	selectedSpecialCharacters := []rune{}
	for j := 0; j < iterations; j++ {
		selectedSpecialCharacters = append(selectedSpecialCharacters, specialChars[getRandom(0, (len(specialChars)-1))])
	}

	for k := 0; k < iterations; k++ {
		chars[k] = selectedSpecialCharacters[k]
	}

	return string(chars)
}
