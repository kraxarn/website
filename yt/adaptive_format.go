package yt

type AdaptiveFormat struct {
	Url              string `json:"url"`
	MimeType         string `json:"mime_type"`
	Bitrate          int    `json:"bitrate"`
	Width            int    `json:"width,omitempty"`
	Height           int    `json:"height,omitempty"`
	Quality          string `json:"quality"`
	Fps              int    `json:"fps,omitempty"`
	QualityLabel     string `json:"quality_label,omitempty"`
	AverageBitrate   int    `json:"average_bitrate,omitempty"`
	ApproxDurationMs string `json:"duration,omitempty"`
	AudioSampleRate  string `json:"audio_sample_rate,omitempty"`
}
