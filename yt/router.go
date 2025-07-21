package yt

import (
	"errors"
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

func Route(app *echo.Echo) {
	app.GET("/yt/search", func(ctx echo.Context) error {
		q := ctx.QueryParam("q")

		if len(q) < 3 {
			return errors.New("query required")
		}

		results, err := search(q)
		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, results)
	})

	app.GET("/yt/info/:id", func(ctx echo.Context) error {
		info, err := Info(ctx.Param("id"))

		if err != nil {
			return err
		}

		return ctx.JSON(http.StatusOK, info)
	})

	app.GET("/yt/audio/:id", func(ctx echo.Context) error {
		info, err := Info(ctx.Param("id"))
		if err != nil {
			return err
		}

		var data []byte
		data, err = helper.Get(info.Audio.Url)
		if err != nil {
			return err
		}

		return ctx.Blob(http.StatusOK, "audio/opus", data)
	})
}
