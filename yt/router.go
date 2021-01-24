package yt

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/common"
	"net/http"
)

func Route(router *gin.Engine) {
	router.GET("/yt/search", func(context *gin.Context) {
		q := context.Query("q")

		if len(q) < 3 {
			context.JSON(http.StatusOK, map[string]interface{}{
				"error": "query required",
			})
			return
		}

		results, err := search(q)
		if err != nil {
			context.JSON(http.StatusOK, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		context.JSON(http.StatusOK, results)
	})

	router.GET("/yt/info/:id", func(context *gin.Context) {
		info, err := Info(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusOK, common.NewError(err))
			return
		}

		context.JSON(http.StatusOK, info)
	})

	router.GET("/yt/audio/:id", func(context *gin.Context) {
		info, err := Info(context.Param("id"))
		var data []byte

		if err != nil {
			fmt.Println("failed to fetch audio info:", err)
		} else {
			data, err = common.Get(info.Audio.Url)
		}
		if err != nil {
			fmt.Println("failed to download audio:", err)
		}

		context.Data(http.StatusOK, "audio/opus", data)
	})
}
