package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"log"
)

// Ofuscate password - hashes password using SHA-256
func hashPassword(password string) (string, error) {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]), nil
}

// Check if password matches hash
func checkPasswordHash(password, hash string) bool {
	hashedPassword, _ := hashPassword(password)
	return hashedPassword == hash
}

func generateRandomToken() (string, error) {
	token := make([]byte, 32)
	_, err := rand.Read(token)
	if err != nil {
		log.Fatalf("Error generating random token: %v", err)
		return "", err
	}
	return base64.URLEncoding.EncodeToString(token), nil
}
