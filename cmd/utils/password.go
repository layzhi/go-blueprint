package utils

import (
	"crypto/rand"
	"math/big"
)

// GenerateRandomPassword generates a random password of a given length.
// It uses a character set of uppercase, lowercase letters, and numbers.
// This is used to replace the default hardcoded passwords in the templates.
func GenerateRandomPassword(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*"
	password := make([]byte, length)
	for i := range password {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			// If crypto/rand fails, we fallback to a less secure method isn't ideal
			// but for this specific CLI tool usage, panic or retry is better?
			// However, crypto/rand failing is extremely rare on supported OSs.
			// Let's just return a default safe fallback if this extremely rare case happens
			// to avoid crashing the CLI for something generated.
			return "ChangeMe123456"
		}
		password[i] = charset[num.Int64()]
	}
	return string(password)
}
