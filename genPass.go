package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func genPass(password, level string) string {
	minLength := 6
	lowercase := "abcdefghijklmnopqrstuvwxyz"
	digits := "0123456789"
	specialChars := "!@#$%^&*()_+{}[]|:;<>,.?/~"

	addRandomChar := func(chars string) {
		randomIndex := rand.Intn(len(chars))
		password += string(chars[randomIndex])
	}

	addUppercase := func() {
		numUppercase := rand.Intn(len(password)-1) + 1

		for i := 0; i < numUppercase; i++ {
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
			password = password[:randomIndex] + string('0'+randomDigit) + password[randomIndex:]
		}
	}

	addRandomSpecialChars := func() {
		numSpecialChars := rand.Intn(len(password))

		for i := 0; i < numSpecialChars; i++ {
			randomIndex := rand.Intn(len(password)-1) + 1
			randomSpecialChar := rand.Intn(len(specialChars))
			password = password[:randomIndex] + string(randomSpecialChar) + password[randomIndex:]
		}
	}

	for len(password) < minLength {
		addRandomChar(lowercase)
	}

	switch strings.ToLower(level) {
	case "low", "med", "strong":
		addUppercase()
	default:
		fmt.Println("Invalid level. Please enter 'low', 'med', or 'strong'.")
		return ""
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

// func main() {
// 	var password, level string
// 	fmt.Print("Enter password: ")
// 	fmt.Scan(&password)

// 	for {
// 		fmt.Print("Enter level (low, med, strong): ")
// 		fmt.Scan(&level)

// 		if strings.ToLower(level) == "low" || strings.ToLower(level) == "med" || strings.ToLower(level) == "strong" {
// 			break
// 		} else {
// 			fmt.Println("Invalid level. Please enter 'low', 'med', or 'strong'.")
// 		}
// 	}

// 	result := genPass(password, level)
// 	if result != "" {
// 		fmt.Println(result)
// 	}
// }
