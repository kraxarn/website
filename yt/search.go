package yt

import (
	"encoding/json"
	"fmt"
	"net/http"
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

func search(query string) ([]SearchResult, error) {
	result, err := http.Get(fmt.Sprintf("%s/api/v1/search?q=%s", invidiousPath, query))
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	var responses []SearchResponse
	err = json.NewDecoder(result.Body).Decode(&responses)
	if err != nil {
		return nil, err
	}

	var results []SearchResult
	for _, response := range responses {
		results = append(results, SearchResult{
			Description: response.Description,
			Id:          response.VideoId,
			Thumbnail:   response.VideoThumbnails[0].Url,
			Title:       response.Title,
			Author:      response.Author,
		})
	}

	return results, nil
}
