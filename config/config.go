package config

import (
	"os"
)

func Dev() bool {
	val, ok := os.LookupEnv("DEV")
	return ok && val == "1"
}
