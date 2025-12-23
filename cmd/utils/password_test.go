package utils

import (
	"testing"
)

func TestGenerateRandomPassword(t *testing.T) {
	length := 16
	password := GenerateRandomPassword(length)

	if len(password) != length {
		t.Errorf("expected password length %d, got %d", length, len(password))
	}

	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	for _, char := range password {
		found := false
		for _, c := range charset {
			if byte(char) == byte(c) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("password contains invalid character: %c", char)
		}
	}

	// Test uniqueness (basic)
	password2 := GenerateRandomPassword(length)
	if password == password2 {
		t.Errorf("generated duplicate passwords, which is extremely unlikely")
	}
}
