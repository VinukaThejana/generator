package utils

// Errors is a struct that contains all the error messages
// that is used throught the application
type Errors struct{}

// FailedToGenerate is when the password failes to generate
// due to various reasons that cannot be properly identified
func (e Errors) FailedToGenerate() string {
	return "Failed to generate password"
}

// ToShort is called when the user inputs an invalid length for the
// password is provided by the user
func (e Errors) ToShort() string {
	return "Length of the password is too short for password generation"
}
