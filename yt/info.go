package yt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type VideoInfo struct {
	Title       string         `json:"title"`
	Thumbnail   string         `json:"thumbnail"`
	Video       AdaptiveFormat `json:"video"`
	Audio       AdaptiveFormat `json:"audio"`
	Description string         `json:"description"`
}

type AdaptiveFormat struct {
	Type       string `json:"type"`
	Url        string `json:"url"`
	Resolution string `json:"resolution,omitempty"`
	Bitrate    string `json:"bitrate"`
}

type VideoInfoResponse struct {
	Title           string
	VideoThumbnails []VideoThumbnailResponse
	Description     string
	AdaptiveFormats []AdaptiveFormat
}

func (i *VideoInfoResponse) videoUrl() AdaptiveFormat {
	var bestFormat AdaptiveFormat
	bestRes := 0

	for _, format := range i.AdaptiveFormats {
		if len(format.Resolution) > 0 {
			if res, err := strconv.Atoi(format.Resolution[:len(format.Resolution)-1]); err == nil {
				if res > bestRes {
					bestFormat = format
					bestRes = res
				}
			}
		}
	}

	return bestFormat
}

func (i *VideoInfoResponse) audioUrl() AdaptiveFormat {
	var bestFormat AdaptiveFormat
	bestBitrate := 0

	for _, format := range i.AdaptiveFormats {
		if strings.HasPrefix(format.Type, "audio") {
			if bitrate, err := strconv.Atoi(format.Bitrate); err == nil {
				if bitrate > bestBitrate {
					bestFormat = format
					bestBitrate = bitrate
				}
			}
		}
	}

	return bestFormat
}

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
