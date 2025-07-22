package group

import (
	"crypto/sha3"
	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/repo"
	"github.com/labstack/echo/v4"
	"net"
	"net/http"
)

func RegisterAdmin(app *echo.Echo) {
	group := app.Group("/admin")

	group.POST("/new", newUser)
}

func newUser(ctx echo.Context) error {
	ip := net.ParseIP(ctx.RealIP())
	if ip == nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid ip")
	}

	if !ip.IsLoopback() {
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")
	if username == "" || password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "required fields missing")
	}

	passwordHashed := sha3.Sum512([]byte(password))

	conn, err := db.Acquire()
	if err != nil {
		return err
	}

	defer conn.Release()
	users := repo.NewUsers(conn)

	var userId db.Id
	userId, err = users.Insert(username, passwordHashed[:], repo.UserDefault)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"id":    userId,
		"flags": repo.UserDefault,
	})
}
