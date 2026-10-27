package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"time"
)

func GetCurrentTimestampMillis() int64 {
	return time.Now().UnixMilli()
}

func GenerateRandomHash(length int) (string, error) {
	const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	hash := make([]byte, length)

	for i := range hash {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("generate random hash: %w", err)
		}

		hash[i] = charset[n.Int64()]
	}

	return string(hash), nil
}

func GenerateUserID() (string, error) {
	timestamp := GetCurrentTimestampMillis()

	randomHash, err := GenerateRandomHash(10)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d%s", timestamp, randomHash), nil
}

f