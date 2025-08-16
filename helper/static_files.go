package helper

import (
	"fmt"
	"github.com/kraxarn/website/config"
	"os"
	"path"
)

var fileVersions map[string]int64

func staticFileVersion(url string) string {
	if config.Dev() {
		return url
	}

	var version int64

	if fileVersions == nil {
		fileVersions = map[string]int64{}
	}

	var ok bool
	if version, ok = fileVersions[url]; !ok {
		info, err := os.Stat(path.Join("static", url))
		if err == nil {
			version = info.ModTime().UnixMilli()
			fileVersions[url] = version
		} else {
			version = 1

		}
	}

	return fmt.Sprintf("%s?v=%d", url, version)
}
