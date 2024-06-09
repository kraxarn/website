package hitster

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
	"time"
)

type HitsterRepository struct {
	conn *pgx.Conn
	ctx  context.Context
}

func connect() (*HitsterRepository, error) {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres:///hitster")
	if err != nil {
		return nil, err
	}

	repo := &HitsterRepository{
		conn: conn,
		ctx:  ctx,
	}

	return repo, nil
}

func (repo *HitsterRepository) Close() error {
	return repo.conn.Close(context.Background())
}

func (repo *HitsterRepository) getAllPlaylists(playlists chan Playlist) error {
	rows, err := repo.conn.Query(repo.ctx,
		"select id, name from playlists",
	)

	if err != nil {
		return err
	}

	defer rows.Close()
	defer close(playlists)

	for rows.Next() {
		playlist := Playlist{}

		err = rows.Scan(&playlist.Id, &playlist.Name)
		if err != nil {
			return err
		}

		playlists <- playlist
	}

	return nil
}

func (repo *HitsterRepository) getRandomPlaylistTracks(playlistId string, limit int, tracks chan Track) error {
	rows, err := repo.conn.Query(repo.ctx, `
        select tracks.id, tracks.name, artists.name, albums.release_date, tracks.preview_url
        from tracks
             join playlist_tracks on tracks.id = playlist_tracks.track_id
             join track_artists on tracks.id = track_artists.track_id
             join artists on track_artists.artist_id = artists.id
             join album_tracks on tracks.id = album_tracks.track_id
             join albums on album_tracks.album_id = albums.id
        where length(preview_url) > 0
          and playlist_tracks.playlist_id = $1
        order by random()
        limit $2
    `, playlistId, limit)

	if err != nil {
		return err
	}

	defer rows.Close()
	defer close(tracks)

	for rows.Next() {
		track := Track{}

		err = rows.Scan(&track.Id, &track.Name, &track.ArtistName, &track.AlbumReleaseDate, &track.PreviewUrl)
		if err != nil {
			return err
		}

		tracks <- track
	}

	return nil
}

func (repo *HitsterRepository) getMinReleaseDate() (time.Time, error) {
	rows, err := repo.conn.Query(repo.ctx,
		"select release_date from albums order by release_date limit 1",
	)

	if err != nil {
		return time.Time{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var date time.Time
		err = rows.Scan(&date)

		if err != nil {
			return time.Time{}, err
		}

		return date, nil
	}

	return time.Time{}, errors.New("no rows")
}

func (repo *HitsterRepository) getMaxReleaseDate() (time.Time, error) {
	rows, err := repo.conn.Query(repo.ctx,
		"select release_date from albums order by release_date desc limit 1",
	)

	if err != nil {
		return time.Time{}, err
	}

	defer rows.Close()

	for rows.Next() {
		var date time.Time
		err = rows.Scan(&date)

		if err != nil {
			return time.Time{}, err
		}

		return date, nil
	}

	return time.Time{}, errors.New("no rows")
}
