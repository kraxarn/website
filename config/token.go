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

func path() string {
	return GetPath("key.bin")
}

func loadKey() ([]byte, error) {
	key, err := os.ReadFile(path())
	if err == nil && len(key) == keySize {
		return key, nil
	}

	key = make([]byte, keySize)
	_, _ = rand.Read(key)

	err = os.WriteFile(path(), key, 0644)
	if err != nil {
		return nil, err
	}

	return key, nil
}

func (t Token) Key() []byte {
	return t.key
}
