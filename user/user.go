package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/google/uuid"
	"github.com/kraxarn/website/config"
	"math/rand"
)

type User struct {
	Id     uuid.UUID
	Name   string
	Avatar uint32
}

func NewUser() *User {
	avatar := AvatarValues[rand.Intn(len(AvatarValues))]

	user := new(User)
	user.Id = uuid.New()
	user.Name = fmt.Sprintf("Anonymous %s", avatar.Name)
	user.Avatar = avatar.Id

	return user
}

func NewUserFromToken(tokenString string, tokenKey *config.Token) (*User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return tokenKey.GetKey(), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := new(User)

		if userName, ok := claims["Name"].(string); ok {
			user.Name = userName
		}
		if userAvatar, ok := claims["Avatar"].(float64); ok {
			user.Avatar = uint32(userAvatar)
		}
		if userId, ok := claims["Id"].(string); ok {
			if userUuid, err := uuid.Parse(userId); err == nil {
				user.Id = userUuid
			}
		}

		return user, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func (user *User) Valid(*jwt.ValidationHelper) error {
	return nil
}

func (user *User) AvatarName() string {
	return AvatarName(user.Avatar)
}

func (user *User) AvatarPath() string {
	return fmt.Sprintf("/img/avatar/%x.svg", user.Avatar)
}

func (user *User) ToToken(token *config.Token) (string, error) {
	return jwt.NewWithClaims(jwt.SigningMethodHS256, user).SignedString(token.GetKey())
}

func (user *User) ToJson() map[string]string {
	return map[string]string{
		"id":     user.Id.String(),
		"name":   user.Name,
		"avatar": fmt.Sprintf("%x", user.Avatar),
	}
}
