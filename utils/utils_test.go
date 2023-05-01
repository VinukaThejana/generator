package utils

import (
	"fmt"
	"log"
	"testing"
)

// TestGetRandomInt is a function that is used for testing the getRandom function
// to test wether the random number is returned within the given range
func TestGetRandomInt(t *testing.T) {
	randomNum := getRandomInt(0, 10)
	fmt.Println("The random number is : ", randomNum)

	if randomNum < 0 || randomNum > 10 {
		t.Errorf("getRandom() returned a number outside the expected range: %d", randomNum)
	}
}

// TestGetRandomLetter is a function that is used to test the function that is used to
// generate a random letter
func TestGetRandomLetter(t *testing.T) {
	randomLetter := getRandomLetter()
	fmt.Println("The random letter is : ", randomLetter)
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

// TestGenerate is a function that is used ot check wether passowrds are generated properly
// with the given conditions
func TestGenerate(t *testing.T) {
	p := Passwords{}

	tests := []int{0, 12, 48, 47, 32, 33, 99, 120}
	for _, test := range tests {
		var password string
		expectedLength := 32

		if test == 0 {
			pswd, err := p.Generate(nil)
			if err != nil {
				log.Println(err)
				t.Error("Failed to generate password")
				return
			}
			password = *pswd
		} else {
			expectedLength = test
			pswd, err := p.Generate(&test)
			if err != nil {
				log.Println(err)
				t.Error("Failed to generate the password")
				return
			}
			password = *pswd
		}

		if len(password) != expectedLength {
			t.Error("The password does not contain enough length")
			return
		}
	}
}
