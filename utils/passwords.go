package utils

import (
	"encoding/base64"
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

// GenerateAlpha is a function is a function that is used to genrate a password
// that only contains alphabetical characters only
func (p Passwords) GenerateAlpha(length *int) (*string, error) {
	passwordLength := 32
	if length != nil {
		passwordLength = *length
		if passwordLength <= 2 {
			return nil, fmt.Errorf("Password length is not sufficent for password generation")
		}
	}

	var password string
	for i := 0; i < passwordLength; i++ {
		password += getRandomLetter()
	}

	return &password, nil
}

// GenerateNumeric is a function that is used to generate  a passwrd that is entirely
// numeric
func (p Passwords) GenerateNumeric(length *int) (*string, error) {
	passwordLength := 32
	if length != nil {
		passwordLength = *length
		if passwordLength <= 2 {
			return nil, fmt.Errorf("Password length is not sufficent for password generation")
		}
	}

	var password string
	for i := 0; i < passwordLength; i++ {
		password += fmt.Sprint(getRandomInt(0, 9))
	}

	return &password, nil
}

// GenerateSpecialChars is a function that is used to generate passwords that only
// contains special characters
func (p Passwords) GenerateSpecialChars(length *int) (*string, error) {
	passwordLength := 32
	if length != nil {
		passwordLength = *length
		if passwordLength <= 2 {
			return nil, fmt.Errorf("Password length is not sufficent for password generation")
		}
	}

	var password string
	for i := 0; i < passwordLength; i++ {
		password += string(getRandomSpecialCharacters())
	}

	return &password, nil
}

// GenerateBase64 is a function that is used ot generate passwords in
// base 64 format
func (p Passwords) GenerateBase64(length *int) (*string, error) {
	password, err := p.Generate(length)
	if err != nil {
		return nil, err
	}

	encoded := base64.StdEncoding.EncodeToString([]byte(*password))
	return &encoded, nil
}
