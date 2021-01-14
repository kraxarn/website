package yt

import (
	"fmt"
	"github.com/kraxarn/website/common"
	"net/url"
	"regexp"
	"strings"
)

// Can also be dynamically fetched with getJsBase
const baseJsUrl string = "https://www.youtube.com/s/player/1a1b48e5/player_ias.vflset/sv_SE/base.js"

var baseJs string

func getJsBase(videoId string) (string, error) {
	body, err := common.Get(fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId))
	if err != nil {
		return "", err
	}

	expr, err := regexp.Compile("\"assets\":.+?\"js\":\\s*\"([^\"]+)\"")
	if err != nil {
		return "", err
	}

	js := expr.FindString(string(body))
	js = strings.Replace(js[strings.LastIndex(js, ":"):], "\\/", "/", -1)
	js = js[2 : len(js)-1]
	if len(js) == 0 {
		err = fmt.Errorf("no player js found")
	}

	return js, err
}

func decrypt(cipher, videoId string) (string, error) {
	path, err := url.QueryUnescape(cipher)
	if err != nil {
		return "", err
	}

	sig, err := between(path, "&sig=", "&")
	if err != nil {
		return "", fmt.Errorf("no signature found (%v)", err)
	}

	return "", nil
}
