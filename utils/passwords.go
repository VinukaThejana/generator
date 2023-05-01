package utils

import (
	"fmt"
	"math"
	"os/exec"
	"strings"
)

// Passwords struct contains all the possible ways that a password
// can be generated with this application
type Passwords struct{}

// Generate is used to generate a passsword that has a default length
// of 32 (unless changed) and contains letters (upper and lower)
// , symbols and numeric characters
func (p Passwords) Generate(length *int) (*string, error) {
	isOdd := false
	passwordLength := 32

	if length != nil {
		passwordLength = *length
		if passwordLength%2 != 0 {
			isOdd = true
		}
	}
	if passwordLength <= 1 {
		return nil, fmt.Errorf("Password length is not sufficent for password generation")
	}

	passwordLength = int(math.Ceil(float64(passwordLength)/2.0) * 2.0)

	out, err := exec.Command("openssl", "rand", "-hex", fmt.Sprint(passwordLength/2)).Output()
	if err != nil {
		return nil, err
	}

	password := strings.TrimSuffix(string(out[:len(out)-1]), "\n")
	cmd := exec.Command("tr", "[:lower:]", "[:upper:]")
	cmd.Stdin = strings.NewReader(password)
	out, err = cmd.Output()
	if err != nil {
		return nil, err
	}

	password = strings.TrimSuffix(string(out[:len(out)-1]), "\n")
	if !isOdd {
		password = fmt.Sprintf("%s%s", password, getRandomLetter())
	}

	password = randomizeWithSymbols(password)
	return &password, nil
}
