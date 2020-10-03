package yt

type AdaptiveFormat struct {
	Type       string `json:"type"`
	Url        string `json:"url"`
	Resolution string `json:"resolution,omitempty"`
	Bitrate    string `json:"bitrate"`
}
