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
		"0CE8F36410AD9D99C2B6E1A6E73D8619CCA3E72A66C849DD1ADF58CEF3AC2137FD616A267205C9358C4D1B98B46D7AB6BDABF27A4A8F63F69684FE0C0D8855BCB8FFA0A0D4CE0E4F6BDF28AF83DAC8EA386D2189C564CAD226CBD81E7E173A679F172F83946A7557C5F2FE10D046C0FC48F3DF800942606DBC5289394401DB5B267C50004A5D0B7377D289028840D5800699DC5BD8D7675B9ED7D7C106CF29EDEB325BFAB7A9DE7A552AC6B6F4B6505F9A53D9987B53642502A08BD48E8463E9FD0E0289ACCDBE5197F7C6CE21AF00462ED29EF3AA1308AEF91A8522AC7D0FD1D219579EF23A9EE867A321F3B4BAB440B8A44A59A398C3A1F07739263E6E1CB307D47F19DF7A7223AE2D950269745F3B0B14741BCA177C2D7AEE1101FB6EA52FC5FD1718FF7AE3EED4ECD87345D231C4A764FFC34751508D99D2E20F50D541225EBD5D297EA588BA0568B3F072DA071C7A6F4E93A4369560C83DCDF61493885A62DEF4A7D01FF94B6FB5E59505802852CC5D07CCDDCB444F49C46D2DE7ABC88AEE90B950D84DDDC2245B82E46334B0DA9A27DCFF800733A67AA8E947A2C5AEC30954E4F884ACFEBDDD13E5D2E42E3926FD8E36FAC902EAF15C3EE0464597014BF94B9A7918C693666B802C1EDEFE4F0448BB0D6ABB9863A67D80FC2B02778D3FA02B961819ADCBC0DF984F55B2370DB95AFC14Fx",
	}

	for n, test := range tests {
		value := randomizeWithSymbols(test)
		fmt.Printf("\t%d). Input :  %s\n\t    Output : %s\n\n", n+1, test, value)
	}
}

// TestGenerate is a function that is used ot check wether passowrds are generated properly
// with the given conditions
func TestGenerate(t *testing.T) {
	g := Generator{}

	tests := []int{0, 12, 48, 47, 32, 33, 99, 120}
	for _, test := range tests {
		var password string
		expectedLength := 32

		if test == 0 {
			pswd, err := g.Generate(nil)
			if err != nil {
				log.Println(err)
				t.Error("Failed to generate password")
				return
			}
			password = *pswd
		} else {
			expectedLength = test
			pswd, err := g.Generate(&test)
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
