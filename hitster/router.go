package hitster

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleError(ctx *gin.Context, err error) {
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
}

func Route(router *gin.Engine) {
	router.GET("/hitster", func(ctx *gin.Context) {
		repo, err := connect()
		handleError(ctx, err)

		defer func(repo *HitsterRepository) {
			handleError(ctx, repo.Close())
		}(repo)

		minDate, err := repo.getMinReleaseDate()
		handleError(ctx, err)

		maxDate, err := repo.getMaxReleaseDate()
		handleError(ctx, err)

		playlists := make(chan Playlist)
		go func() {
			handleError(ctx, repo.getAllPlaylists(playlists))
		}()

		if ctx.IsAborted() {
			return
		}

		ctx.HTML(http.StatusOK, "hitster.gohtml", gin.H{
			"playlists": playlists,
			"minDate":   minDate,
			"maxDate":   maxDate,
		})
	})

	router.GET("/hitster/:id", func(ctx *gin.Context) {
		repo, err := connect()
		handleError(ctx, err)

		defer func(repo *HitsterRepository) {
			handleError(ctx, repo.Close())
		}(repo)

		playlistId := ctx.Param("id")

		tracks := make(chan Track)
		go func() {
			err = repo.getRandomPlaylistTracks(playlistId, 10, tracks)
			handleError(ctx, err)
		}()

		var slice []Track
		for track := range tracks {
			slice = append(slice, track)
		}

		ctx.JSON(http.StatusOK, slice)
	})
}
