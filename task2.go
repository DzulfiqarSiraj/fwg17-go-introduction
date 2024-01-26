package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func PassGenerator(password string, level string) string {
	minPassLen := 6
	lowerCaseChars := "abcdefghijklmnopqrstuvwxyz"
	digits := "1234567890"
	specialChars := "!@#$%^&*()_+{}|[];:'<>?,./-=~"
	level = strings.ToLower(level)

	genRandomChar := func(chars string) {
		index := rand.Intn(len(chars))
		password = password + string(chars[index])
	}

	genUpperCase := func() {
		upperCaseRange := rand.Intn(len(password))

		for i := 0; i <= upperCaseRange; i++ {
			index := rand.Intn(len(password))
			upperCaseChar := strings.ToUpper(string(password[index]))
			password = password[:index] + upperCaseChar + password[index+1:]
		}
	}

	genRandomDigits := func() {
		digitRange := rand.Intn(len(password))

		for i := 0; i < digitRange; i++ {
			index := rand.Intn(len(password))
			randDigit := rand.Intn(len(digits))
			password = password[:index] + string(digits[randDigit]) + password[index:]
		}
	}

	genRandomSpecialChars := func() {
		specialCharsRange := rand.Intn(len(password))

		for i := 0; i < specialCharsRange; i++ {
			index := rand.Intn(len(password))
			randSpecialChars := rand.Intn(len(specialChars))
			password = password[:index] + string(specialChars[randSpecialChars]) + password[index:]
		}
	}

	for len(password) < minPassLen {
		genRandomChar(lowerCaseChars)
	}

	if level == "low" || level == "med" || level == "strong" {
		genUpperCase()
	} else {
		fmt.Println("Invalid input. Input must be (low / med / strong)")
		return ""
	}

	if level == "med" || level == "strong" {
		genRandomDigits()
	}

	if level == "strong" {
		genRandomSpecialChars()
	}

	return password
}
