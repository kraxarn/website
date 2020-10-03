package yt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func info(videoId string) (VideoInfo, error) {
	result, err := http.Get(fmt.Sprintf("%s/api/v1/videos/%s", invidiousPath, videoId))
	if err != nil {
		return VideoInfo{}, err
	}

	defer func() {
		_ = result.Body.Close()
	}()

	var response VideoInfoResponse
	err = json.NewDecoder(result.Body).Decode(&response)
	if err != nil {
		return VideoInfo{}, err
	}

	return VideoInfo{
		Title:       response.Title,
		Thumbnail:   response.VideoThumbnails[0].Url,
		Video:       response.videoUrl(),
		Audio:       response.audioUrl(),
		Description: response.Description,
	}, nil
}
