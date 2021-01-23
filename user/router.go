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
			"Gets info about a user from token",
		},
		{
			http.MethodGet, "whoami", manager.whoAmI,
			"Gets info about the current user",
		},
	}

	for _, handler := range manager.handlers {
		router.Handle(handler.httpMethod, handler.relativePath(), handler.handler)
	}
}
