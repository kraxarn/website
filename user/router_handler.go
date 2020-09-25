package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type RouterHandler struct {
	httpMethod  string
	path        string
	handler     gin.HandlerFunc
	description string
}

func (handler *RouterHandler) relativePath() string {
	return fmt.Sprintf("/user/%s", handler.path)
}
