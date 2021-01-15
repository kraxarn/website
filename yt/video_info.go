package yt

type VideoInfo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Duration    int    `json:"duration"`
	Thumbnail   string `json:"thumbnail"`
	Audio       Format `json:"audio"`
	Video       Format `json:"video"`
}
