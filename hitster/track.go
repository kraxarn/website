package hitster

import "time"

type Track struct {
	Id               string    `json:"id"`
	Name             string    `json:"name"`
	PreviewUrl       string    `json:"preview_url"`
	ArtistName       string    `json:"artist_name"`
	AlbumReleaseDate time.Time `json:"album_release_date"`
}
