package config

import (
	"crypto/rand"
	"os"
)

const keySize int = 32

type Token struct {
	key []byte
}

func NewToken() (Token, error) {
	key, err := loadKey()
	if err != nil {
		return Token{}, err
	}

	return Token{
		key: key,
	}, nil
}

func keyPath() string {
	return GetPath("key.bin")
}

func loadKey() ([]byte, error) {
	key, err := os.ReadFile(keyPath())
	if err == nil && len(key) == keySize {
		return key, nil
	}

	key = make([]byte, keySize)
	_, _ = rand.Read(key)

	err = os.WriteFile(keyPath(), key, 0644)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (t Token) Key() []byte {
	return t.key
}
