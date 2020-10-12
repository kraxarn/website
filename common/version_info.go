package common

import "strings"

type VersionInfo struct {
	Name    string   `json:"name"`
	Changes []string `json:"changes"`
}

func GetVersions() ([]VersionInfo, error) {
	var releases []interface{}
	if err := GetJson("https://api.github.com/repos/kraxarn/website/releases", &releases); err != nil {
		return nil, err
	}

	var versions []VersionInfo
	for _, release := range releases {
		data := release.(map[string]interface{})

		var changes []string
		for _, change := range strings.Split(data["body"].(string), "\r\n") {
			if len(change) > 2 {
				changes = append(changes, change[2:])
			}
		}

		versions = append(versions, VersionInfo{
			Name:    data["name"].(string),
			Changes: changes,
		})
	}

	return versions, nil
}
