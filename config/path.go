package config

import (
	"fmt"
	"os"
	"path"
)

func GetPath(fileName string) string {
	var err error

	filePath, err := os.UserConfigDir()
	if err == nil {
		filePath = path.Join(filePath, "kraxarn", "website")
		err = os.MkdirAll(filePath, 0755)
		if err == nil {
			return path.Join(filePath, fileName)
		}
	}

	fmt.Printf("error: failed to get path for file \"%s\": %v", fileName, err)
	return fileName
}
