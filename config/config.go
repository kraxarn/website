package config

import (
	"errors"
	"os"
)

func Dev() bool {
	val, ok := os.LookupEnv("DEV")
	return ok && val == "1"
}

func DatabaseUrl() (string, error) {
	val, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		return "", errors.New("no database url set")
	}

	return val, nil
}
