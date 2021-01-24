package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/config"
	"net/http"
)

func Route(router *gin.Engine, token *config.Token) {
	var manager RouterManager
	manager.token = token
	manager.handlers = []RouterHandler{
		{
			http.MethodGet, "api", manager.api,
			"Get all available API calls",
		},
		{
			http.MethodGet, "new", manager.new,
			"Create a new user",
		},
		{
			http.MethodPost, "info", manager.info,
			"Gets info about a user from token, or from stored cookie",
		},
		{
			http.MethodPost, "update", manager.update,
			"Update avatar and/or username of user from token or cookie",
		},
	}

	for _, handler := range manager.handlers {
		router.Handle(handler.httpMethod, handler.relativePath(), handler.handler)
	}
}
