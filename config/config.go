package config

import (
	"fmt"
	"os"
)

func Dev() bool {
	val, ok := os.LookupEnv("DEV")
	return ok && val == "1"
}

func lookupEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("missing key: %s", key)
	}

	return val, nil
}

func DatabaseUrl() (string, error) {
	return lookupEnv("DATABASE_URL")
}

func TeamSpeakUrl() (string, error) {
	return lookupEnv("TEAMSPEAK_URL")
}

func TeamSpeakApiKey() (string, error) {
	return lookupEnv("TEAMSPEAK_API_KEY")
}
