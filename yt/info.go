package yt

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func jsonInfo(videoId string) (map[string]interface{}, error) {
	out, err := exec.Command("youtube-dl", "--print-json", "--simulate",
		fmt.Sprintf("https://www.youtube.com/watch?v=%s", videoId)).Output()
	if err != nil {
		return nil, err
	}

	var data map[string]interface{}
	if err = json.Unmarshal(out, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func Info(videoId string) (VideoInfo, error) {
	data, err := jsonInfo(videoId)
	if err != nil {
		return VideoInfo{}, err
	}

	info := VideoInfo{
		Thumbnail: fmt.Sprintf("https://i.ytimg.com/vi/%s/maxresdefault.jpg", data["id"].(string)),
		Video:     Format{},
		Audio:     Format{},
	}

	if title, ok := data["title"].(string); ok {
		info.Title = title
	}
	if description, ok := data["description"].(string); ok {
		info.Description = description
	}
	if duration, ok := data["duration"].(float64); ok {
		info.Duration = int(duration)
	}

	for _, itf := range data["requested_formats"].([]interface{}) {
		data := itf.(map[string]interface{})
		var format Format

		if url, ok := data["url"].(string); ok {
			format.Url = url
		}
		if width, ok := data["width"].(int64); ok {
			format.Width = int(width)
		}
		if height, ok := data["height"].(int64); ok {
			format.Height = int(height)
		}
		if quality, ok := data["format_note"].(string); ok {
			format.Quality = quality
		}
		if fps, ok := data["fps"].(float64); ok {
			format.Fps = int(fps)
		}
		if bitrate, ok := data["tbr"].(float64); ok {
			format.AverageBitrate = bitrate
		}
		if sampleRate, ok := data["asr"].(float64); ok {
			format.AudioSampleRate = int(sampleRate)
		}

		if audioCodec, ok := data["acodec"].(string); ok && audioCodec != "none" {
			// Audio specific
			format.Codec = audioCodec
			if bitrate, ok := data["abr"].(float64); ok {
				format.Bitrate = int(bitrate)
			}
			info.Audio = format
		} else if videoCodec, ok := data["vcodec"].(string); ok && videoCodec != "none" {
			// Video specific
			format.Codec = videoCodec
			if bitrate, ok := data["vbr"].(float64); ok {
				format.Bitrate = int(bitrate)
			}
			info.Video = format
		}
	}

	return info, nil
}
