package auth

import (
	"crypto/sha256"
)


// HashPassword Hash password passed through the
// parameter
func HashPassword(Password string) (string, error) {
	data := []byte(Password)
	hash := sha256.Sum256(data)
	// returning hash as a string
	return string(hash[:]), nil
}

