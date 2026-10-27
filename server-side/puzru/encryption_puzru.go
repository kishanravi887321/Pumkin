package puzru

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Encrypt encrypts plaintext using your 32-byte AES-256 key.
// Encryption and Key Length Configuration:
//
// 1. User Encryption Key:         10 characters
// 2. Private Server Auth Key:     12 characters
// 3. Master Encryption Key:       10 characters

func Encrypt(plaintext string, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a unique, random nonce for this encryption.
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	// Encrypt the data and prepend the nonce.
	encrypted := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encode the encrypted bytes as a string for storage.
	return base64.RawURLEncoding.EncodeToString(encrypted), nil
}

// Decrypt decrypts data using the same 32-byte AES-256 key.
func Decrypt(encryptedText string, key []byte) (string, error) {
	data, err := base64.RawURLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Ensure the ciphertext contains a nonce and authentication tag.
	if len(data) < gcm.NonceSize()+gcm.Overhead() {
		return "", errors.New("invalid ciphertext")
	}

	// Extract the nonce and ciphertext.
	nonce := data[:gcm.NonceSize()]
	ciphertext := data[gcm.NonceSize():]

	// Decrypt and verify integrity.
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", errors.New("decryption failed: invalid key or ciphertext")
	}

	return string(plaintext), nil
}

// Assume key is your decoded, 32-byte secret key.
