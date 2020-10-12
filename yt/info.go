package yt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func info(videoId string) (VideoInfo, error) {
	result, err := http.Get(fmt.Sprintf("https://www.youtube.com/get_video_info?video_id=%[1]s&eurl=https://youtube.googleapis.com/v/%[1]s", videoId))
	if err != nil {
		return VideoInfo{}, err
	}
	defer func() {
		_ = result.Body.Close()
	}()

	bodyData, err := ioutil.ReadAll(result.Body)
	if err != nil {
		return VideoInfo{}, err
	}
	response, err := between(string(bodyData), "player_response=", "&")
	if err != nil {
		return VideoInfo{}, err
	}
	response, err = url.QueryUnescape(response)
	if err != nil {
		return VideoInfo{}, nil
	}

	var data map[string]interface{}
	err = json.Unmarshal([]byte(response), &data)
	if err != nil {
		return VideoInfo{}, err
	}

	var formats []AdaptiveFormat
	for _, itf := range data["streamingData"].(map[string]interface{})["adaptiveFormats"].([]interface{}) {
		data := itf.(map[string]interface{})

		format := AdaptiveFormat{}
		if urlString, ok := data["url"].(string); ok {
			format.Url = urlString
		} else {
			continue
		}
		if mimeType, ok := data["mimeType"].(string); ok {
			format.MimeType = mimeType
		}
		if bitrate, ok := data["bitrate"].(float64); ok {
			format.Bitrate = int(bitrate)
		}
		if quality, ok := data["quality"].(string); ok {
			format.Quality = quality
		}

		if width, ok := data["width"].(float64); ok {
			format.Width = int(width)
		}
		if height, ok := data["height"].(float64); ok {
			format.Height = int(height)
		}
		if fps, ok := data["fps"].(float64); ok {
			format.Fps = int(fps)
		}
		if qualityLabel, ok := data["qualityLabel"].(string); ok {
			format.QualityLabel = qualityLabel
		}

		if averageBitrate, ok := data["averageBitrate"].(float64); ok {
			format.AverageBitrate = int(averageBitrate)
		}
		if approxDurationMs, ok := data["approxDurationMs"].(string); ok {
			format.ApproxDurationMs = approxDurationMs
		}
		if audioSampleRate, ok := data["audioSampleRate"].(string); ok {
			format.AudioSampleRate = audioSampleRate
		}

		formats = append(formats, format)
	}

	details := data["videoDetails"].(map[string]interface{})
	return VideoInfo{
		Title:       details["title"].(string),
		Thumbnail:   fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", details["videoId"].(string)),
		Video:       bestVideoFormat(formats),
		Audio:       bestAudioFormat(formats),
		Description: details["shortDescription"].(string),
	}, nil
}
