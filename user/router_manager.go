package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/common"
	"github.com/kraxarn/website/config"
	"net/http"
)

type RouterManager struct {
	token    *config.Token
	handlers []RouterHandler
}

func (manager *RouterManager) api(ctx *gin.Context) {
	json := make(map[string]interface{})
	for _, handler := range manager.handlers {
		json[handler.path] = map[string]string{
			"call":        handler.relativePath(),
			"method":      handler.httpMethod,
			"description": handler.description,
		}
	}

	ctx.JSON(http.StatusOK, json)
}

func (manager *RouterManager) new(ctx *gin.Context) {
	user := NewUser()
	token, err := user.ToToken(manager.token)

	var json map[string]interface{}
	if err != nil {
		json = map[string]interface{}{
			"error": fmt.Sprintf("%T: %s", err, err.Error()),
		}
	} else {
		json = map[string]interface{}{
			"token": token,
			"user":  user.ToJson(),
		}
	}

	ctx.JSON(http.StatusOK, json)
}

func (manager *RouterManager) userJson(token string) interface{} {
	user, err := NewUserFromToken(token, manager.token)

	if err != nil {
		return common.NewError(err)
	}
	return map[string]interface{}{
		"user": user.ToJson(),
	}
}

func (manager *RouterManager) info(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, manager.userJson(ctx.PostForm("token")))
}

func (manager *RouterManager) whoAmI(ctx *gin.Context) {
	if token, err := ctx.Cookie("user"); err != nil {
		ctx.JSON(http.StatusOK, common.NewError(err))
	} else {
		ctx.JSON(http.StatusOK, manager.userJson(token))
	}
}
