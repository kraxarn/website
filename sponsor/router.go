package sponsor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Route(router *gin.Engine) {
	manager, err := NewManager()
	if err != nil {
		fmt.Println("failed to instance sponsor manager:", err)
		return
	}

	router.GET("/sponsor/:id", func(context *gin.Context) {
		times, err := manager.GetTimes(context.Param("id"))
		if err != nil {
			context.JSON(http.StatusOK, map[string]interface{}{
				"error": err,
			})
			return
		}
		context.JSON(http.StatusOK, times)
	})
}
