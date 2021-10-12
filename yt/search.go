package yt

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

type SearchResult struct {
	Description string `json:"description"`
	Id          string `json:"id"`
	Thumbnail   string `json:"thumbnail"`
	Title       string `json:"title"`
	Author      string `json:"author"`
}

type VideoThumbnailResponse struct {
	Url string
}

type SearchResponse struct {
	Title           string
	Author          string
	VideoId         string
	VideoThumbnails []VideoThumbnailResponse
	Description     string
}

func between(s, start, end string) (string, error) {
	startIndex := strings.Index(s, start) + len(start)
	if startIndex < 0 {
		return "", fmt.Errorf("start not found in string")
	}

	endIndex := strings.Index(s[startIndex:], end) + startIndex
	if endIndex < startIndex {
		return "", fmt.Errorf("end not found in string")
	}

	return s[startIndex:endIndex], nil
}

func betweenOrEmpty(s, start, end string) string {
	match, err := between(s, start, end)
	if err == nil {
		return match
	}
	return ""
}

func search(query string) ([]SearchResult, error) {
	result, err := http.Get(fmt.Sprintf("https://www.youtube.com/results?search_query=%s",
		url.QueryEscape(query)))
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = result.Body.Close()
	}()

	bodyData, err := io.ReadAll(result.Body)
	if err != nil {
		return nil, err
	}

	body := string(bodyData)
	data, err := between(body, "ytInitialData", ";")
	if err != nil {
		return nil, err
	}

	expr, err := regexp.Compile("videoRenderer.*?serviceEndpoint")
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	for _, match := range expr.FindAllString(data, -1) {
		results = append(results, SearchResult{
			Author:      betweenOrEmpty(match, `"ownerText":{"runs":[{"text":"`, `","`),
			Description: betweenOrEmpty(match, `"descriptionSnippet":{"runs":[{"text":"`, `"}]},`),
			Id:          betweenOrEmpty(match, `"videoId":"`, `"`),
			Thumbnail:   betweenOrEmpty(match, `"thumbnail":{"thumbnails":[{"url":"`, `"`),
			Title:       betweenOrEmpty(match, `"title":{"runs":[{"text":"`, `"}]`),
		})
	}

	if results == nil {
		return nil, fmt.Errorf("no results")
	}
	return results, nil
}
