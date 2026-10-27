package puzru

import (
	"fmt"

	"github.com/kishanravi887321/pumkin/server-side/internal/config"
)

func MakeKey() (string, error) {
	cfg, err := config.Load()
	if err != nil {
		return "", err
	}
	key1, err := CallPrivateServer()
	if err != nil {
		return "", err
	}
	key2 := cfg.MasterEncryptionKey

	key3 := key1 + key2

	fmt.Println("Master Encryption Key:", key3)
	return key1 + key2, nil
}
