package yt

import (
	"github.com/gin-gonic/gin"
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
				"error": err,
			})
			return
		}

		context.JSON(http.StatusOK, results)
	})

	router.GET("/yt/info/:id", func(context *gin.Context) {
		info, err := info(context.Param("id"))

		if err != nil {
			context.JSON(http.StatusOK, map[string]interface{}{
				"error": err,
			})
			return
		}

		context.JSON(http.StatusOK, info)
	})
}
