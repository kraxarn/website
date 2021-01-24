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

	var newUser User
	if err := ctx.BindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusOK, common.NewError(err))
		return
	}
	if len(newUser.Name) <= 0 && newUser.Avatar <= 0 {
		ctx.JSON(http.StatusOK, common.NewError(fmt.Errorf("no changes made")))
		return
	}

	if len(newUser.Name) > 0 {
		currentUser.Name = newUser.Name
	}

	if newUser.Avatar > 0 && AvatarExists(newUser.Avatar) {
		currentUser.Avatar = newUser.Avatar
	}

	// Only refresh cookie if we had one before
	newToken, _ := currentUser.ToToken(manager.token)
	if _, err := ctx.Cookie("user"); len(newToken) > 0 && err == nil {
		currentUser.RefreshWithToken(ctx, newToken)
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"token": newToken,
		"user":  currentUser,
	})
}
