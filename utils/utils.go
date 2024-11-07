package utils

import (
	"math/rand"

	"github.com/atotto/clipboard"
)

func GeneratePassword(length int, includeDigits bool, includeSymbols bool) string {
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	if includeDigits {
		charset += "0123456789"
	}
	if includeSymbols {
		charset += "!@#$%^&*()-_=+,.?/:;{}[]`~"
	}

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	return string(password)
}

func WriteToClipboard(text string) error {
	return clipboard.WriteAll(text)
}
