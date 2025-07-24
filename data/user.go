package data

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/kraxarn/website/db"
	"github.com/labstack/echo/v4"
	"strconv"
)

type UserFlags uint

const (
	UserFlagsNone   UserFlags = 0
	UserFlagsLogin  UserFlags = 1
	UserFlagsEditor UserFlags = 2
)

type UserClaims struct {
	claims jwt.MapClaims
}

func ParseUserClaims(ctx echo.Context) (UserClaims, error) {
	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return UserClaims{}, errors.New("no token")
	}

	return UserClaims{
		claims: token.Claims.(jwt.MapClaims),
	}, nil
}

func (c UserClaims) UserId() (db.Id, error) {
	sub, err := c.claims.GetSubject()
	if err != nil {
		return 0, err
	}

	var val int64
	val, err = strconv.ParseInt(sub, 10, 64)
	if err != nil {
		return 0, err
	}

	return db.Id(val), nil
}

func (c UserClaims) UserFlags() (UserFlags, error) {
	flg, ok := c.claims["flg"].(string)
	if !ok {
		return UserFlagsNone, errors.New("no such key")
	}

	val, err := strconv.ParseUint(flg, 10, 64)
	if err != nil {
		return UserFlagsNone, err
	}

	return UserFlags(val), nil
}
