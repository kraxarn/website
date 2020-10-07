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

func bestVideoFormat(formats []AdaptiveFormat) AdaptiveFormat {
	var bestFormat AdaptiveFormat
	bestRes := 0

	for _, format := range formats {
		if len(format.QualityLabel) > 0 {
			if res, err := strconv.Atoi(format.QualityLabel[:len(format.QualityLabel)-1]); err == nil {
				if res > bestRes {
					bestFormat = format
					bestRes = res
				}
			}
		}
	}

	return bestFormat
}

func bestAudioFormat(formats []AdaptiveFormat) AdaptiveFormat {
	var bestFormat AdaptiveFormat
	bestBitrate := 0

	for _, format := range formats {
		if strings.HasPrefix(format.MimeType, "audio") {
			if format.Bitrate > bestBitrate {
				bestFormat = format
				bestBitrate = format.Bitrate
			}
		}
	}

	return bestFormat
}
