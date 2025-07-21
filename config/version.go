package config

import (
	"fmt"
)

const (
	VersionMajor uint8 = 14
	VersionMinor uint8 = 0
	VersionPatch uint8 = 0
)

func Version() string {
	return fmt.Sprintf("v%d.%d.%d", VersionMajor, VersionMinor, VersionPatch)
}
