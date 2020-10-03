package yt

import (
	"strconv"
	"strings"
)

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
