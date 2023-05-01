package utils

import (
	"fmt"
	"testing"
)

// TestGetRandom is a function that is used for testing the getRandom function
// to test wether the random number is returned within the given range
func TestGetRandom(t *testing.T) {
	randomNum := getRandom(0, 10)
	fmt.Println("The random number is : ", randomNum)

	if randomNum < 0 || randomNum > 10 {
		t.Errorf("getRandom() returned a number outside the expected range: %d", randomNum)
	}
}

// TestRandomizeWithSymbols is a function that is used to test strings inorder
// to check wether the strings will be properly randomized
func TestRandomizeWithSymbols(t *testing.T) {
	tests := []string{
		"g/abjhYbsV5GZcLzTqAZ1gZOSnHFfMsKa38GAm2KsdY=",
		"BezEbiGuFAo7LFuq8E4Q8aoF6fc0K7FkVLQyNCLTZnW+8zdrr8AGmHf/X18flXG0",
		"pnmRZuFA2SC6gWnqBjXzKrd+/mZ52OWdHQiS0J3tFysnCJmqZ/7YeSoxBL3Wkij26LtE6CgWUypVl+1pXVsCWw==",
		"ZxFSnQWaQsFEfDqMAyyAjuvxZhXTFARh",
		"wzeuqGnECJQESNGWQsBfQDeDnpNgqWyesvUPACAKZgDxgRFpHN",
	}

	for n, test := range tests {
		value := randomizeWithSymbols(test)
		fmt.Printf("\t%d). Input :  %s\n\t    Output : %s\n\n", n+1, test, value)
	}
}
