package config

import (
	"fmt"
	"math"
	"math/rand"
	"os"
)

type Token struct {
	key  [32]byte
	path string
}

func NewToken() Token {
	t := Token{}
	t.load()
	return t
}

func getPath() string {
	return GetPath("key.bin")
}

func (token *Token) load() {
	token.path = getPath()

	if data, err := os.ReadFile(token.path); err != nil && len(data) == 32 {
		for i := range token.key {
			token.key[i] = data[i]
		}
	} else {
		for i := range token.key {
			token.key[i] = uint8(rand.Uint32() % math.MaxUint8)
		}
		if err := os.WriteFile(token.path, token.key[:], 0644); err != nil {
			fmt.Println("error: failed to save key to file:", err)
		}
	}
}

func (token *Token) GetKey() []byte {
	return token.key[:]
}

func (token *Token) GetPath() string {
	return token.path
}
