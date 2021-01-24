package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/common"
	"github.com/kraxarn/website/config"
	"net/http"
	"strconv"
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

func (manager *RouterManager) getUser(ctx *gin.Context) *User {
	// From POST
	token := ctx.PostForm("token")
	// From cookie
	if len(token) <= 0 {
		token, _ = ctx.Cookie("user")
	}
	currentUser, _ := NewUserFromToken(token, manager.token)
	return currentUser
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
	userInfo := manager.getUser(ctx)
	if userInfo == nil {
		ctx.JSON(http.StatusOK, common.NewError(fmt.Errorf("no user found")))
	} else {
		ctx.JSON(http.StatusOK, userInfo)
	}
}

func (manager *RouterManager) update(ctx *gin.Context) {
	currentUser := manager.getUser(ctx)
	if currentUser == nil {
		ctx.JSON(http.StatusOK, common.NewError(fmt.Errorf("no user found")))
		return
	}

	name := ctx.PostForm("name")
	if len(name) > 0 {
		currentUser.Name = name
	}

	avatar, err := strconv.ParseUint(ctx.PostForm("avatar"), 10, 32)
	if err == nil && AvatarExists(uint32(avatar)) {
		currentUser.Avatar = uint32(avatar)
	}

	// Only refresh cookie if we had one before
	var newToken string
	if _, err = ctx.Cookie("user"); err != nil {
		newToken = currentUser.Refresh(ctx, manager.token)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": newToken,
		"user":  currentUser,
	})
}
