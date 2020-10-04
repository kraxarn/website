package yt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

type SearchResult struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Thumbnail   string `json:"thumbnail"`
	Title       string `json:"title"`
}

type VideoThumbnailResponse struct {
	Url string
}

type SearchResponse struct {
	Title           string
	VideoId         string
	VideoThumbnails []VideoThumbnailResponse
	Description     string
}

// TODO: TESTING ONLY
func Search(query string) ([]SearchResult, error) {
	return search(query)
}

func between(s, start, end string) (string, error) {
	startIndex := strings.Index(s, start) + len(start)
	if startIndex < 0 {
		return "", fmt.Errorf("start not found in string")
	}

	endIndex := strings.Index(s[startIndex:], end) + startIndex
	if endIndex < 0 {
		return "", fmt.Errorf("end not found in string")
	}

	return s[startIndex:endIndex], nil
}

func dig(data map[string]interface{}, parameters []string) (map[string]interface{}, error) {
	endpoint := data

	for i, param := range parameters {
		/*if arr, ok := endpoint[param].([]interface{}); ok {
			endpoint = arr[0].(map[string]interface{})
		}*/
		if newEndpoint, ok := endpoint[param].(map[string]interface{}); ok {
			endpoint = newEndpoint
		} else {
			return endpoint, fmt.Errorf("property %s (index %d) does not exist in map (%v)", param, i, reflect.ValueOf(endpoint).MapKeys())
		}
	}

	return endpoint, nil
}

func search(query string) ([]SearchResult, error) {
	result, err := http.Get(fmt.Sprintf("https://www.youtube.com/results?search_query=%s", query))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = result.Body.Close()
	}()

	bodyData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	body := string(bodyData)
	data, err := between(body, "window[\"ytInitialData\"] = ", ";")
	if err != nil {
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal([]byte(data), &response)
	if err != nil {
		return nil, err
	}

	digResponse, err := dig(response, []string{
		"contents", "twoColumnSearchResultsRenderer", "primaryContents", "sectionListRenderer",
		"contents", "itemSectionRenderer",
	})

	if err != nil {
		return nil, err
	}
	_ = digResponse["contents"].([]interface{})

	return nil, nil
}
