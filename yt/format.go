package yt

type Format struct {
	Url             string  `json:"url"`
	Codec           string  `json:"codec"`
	Bitrate         int     `json:"bitrate"`
	Width           int     `json:"width,omitempty"`
	Height          int     `json:"height,omitempty"`
	Quality         string  `json:"quality"`
	Fps             int     `json:"fps,omitempty"`
	AverageBitrate  float64 `json:"average_bitrate,omitempty"`
	AudioSampleRate int     `json:"audio_sample_rate,omitempty"`
}
