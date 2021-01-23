package watch

import (
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/config"
	"github.com/kraxarn/website/user"
)

type Watch struct {
	token *config.Token
}

func (watch *Watch) getUser(context *gin.Context) *user.User {
	var currentUser *user.User
	token, err := context.Cookie("user")
	if err == nil {
		currentUser, err = user.NewUserFromToken(token, watch.token)
	}

	if currentUser == nil {
		currentUser = user.NewUser()
	}
	if currentUser == nil {
		return nil
	}

	// 1 month
	if token, err := currentUser.ToToken(watch.token); err == nil {
		context.SetCookie("user", token, 2_629_800,
			"/", config.GetDomain(), config.IsSecure(), true)
	}
	return currentUser
}
