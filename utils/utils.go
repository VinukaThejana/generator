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
