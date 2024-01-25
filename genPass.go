package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func genPass(password string, level string) string {
	minLength := 6
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	specialChars := "!@#$%^&*()_+{}[]|:;<>,.?/~"

	addRandomChar := func(chars string) {
		randomIndex := rand.Intn(len(chars))
		password += string(chars[randomIndex])
	}

	randUppercase := func() {
		numUppercase := rand.Intn(len(password)-1) + 1

		for i := 0; i <= numUppercase; i++ {
			randomIndex := rand.Intn(len(password))
			upperChar := strings.ToUpper(string(password[randomIndex]))
			password = password[:randomIndex] + upperChar + password[randomIndex+1:]
		}
	}

	addRandomDigits := func() {
		numDigits := rand.Intn(len(password))

		for i := 0; i < numDigits; i++ {
			randomIndex := rand.Intn(len(password)-1) + 1
			randomDigit := rand.Intn(len(digits))
			password = password[:randomIndex] + string(digits[randomDigit]) + password[randomIndex:]
		}
	}

	addRandomSpecialChars := func() {
		numSpecialChars := rand.Intn(len(password))

		for i := 0; i < numSpecialChars; i++ {
			randomIndex := rand.Intn(len(password)-1) + 1
			randomSpecialChar := rand.Intn(len(specialChars))
			password = password[:randomIndex] + string(specialChars[randomSpecialChar]) + password[randomIndex:]
		}
	}

	for len(password) < minLength {
		addRandomChar(lowercase)
	}

	switch strings.ToLower(level) {
	case "low", "med", "strong":
		randUppercase()
	default:
		fmt.Println("Invalid level. Please enter 'low', 'med', or 'strong'.")
	}

	switch strings.ToLower(level) {
	case "med", "strong":
		addRandomDigits()
	}

	switch strings.ToLower(level) {
	case "strong":
		addRandomSpecialChars()
	}

	return password
}
