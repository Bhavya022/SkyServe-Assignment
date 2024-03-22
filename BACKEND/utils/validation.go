// utils: This folder contains utility functions or helper functions used across the application.

// validation.go: Contains functions for data validation, such as validating user input.
package utils

import (
	"regexp"
)

// ValidateEmail checks if the given email address is valid.
func ValidateEmail(email string) bool {
	// Regular expression for basic email validation
	// This is a simplified regex for demonstration purposes
	// Replace it with a more comprehensive one in a real application
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

// ValidatePassword checks if the given password meets the minimum requirements.
func ValidatePassword(password string) bool {
	// Minimum password length requirement
	minLength := 8
	return len(password) >= minLength
}
