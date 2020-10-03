package yt

type VideoInfo struct {
	Title       string         `json:"title"`
	Thumbnail   string         `json:"thumbnail"`
	Video       AdaptiveFormat `json:"video"`
	Audio       AdaptiveFormat `json:"audio"`
	Description string         `json:"description"`
}
