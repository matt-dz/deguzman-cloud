package api

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"hash"
	"io"
	"os"
)

func generateRandomString(n int) (string, error) {
	// Calculate the number of bytes required to encode `n` characters in base64
	// Base64 encoding expands every 3 bytes into 4 characters.
	byteLength := (n*3 + 3) / 4

	randomBytes := make([]byte, byteLength)
	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}
	encoded := base64.RawURLEncoding.EncodeToString(randomBytes)

	return encoded[:n], nil
}

func HashContents(file io.Reader, hashFunc hash.Hash) (string, error) {
	buffer := make([]byte, 4096) // 4KB buffer
	for {
		n, err := file.Read(buffer)
		if n > 0 {
			if _, writeErr := hashFunc.Write(buffer[:n]); writeErr != nil {
				return "", fmt.Errorf("failed to write to hash function: %v", writeErr)
			}
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			return "", fmt.Errorf("failed to read file: %v", err)
		}
	}

	return fmt.Sprintf("%x", hashFunc.Sum(nil)), nil
}

func HashFile(filePath string, hashFunc hash.Hash) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	return HashContents(file, hashFunc)
}
