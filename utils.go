package main

import (
	"crypto/sha256"
	"encoding/hex"
)

// Ofuscate password - hashes password using SHA-256
func hashPassword(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}
