package group

import (
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/repo"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net"
	"net/http"
)

func RegisterAdmin(app *echo.Echo) {
	group := app.Group("/admin")

	group.GET("", admin)
	group.POST("", login)
	group.POST("/new", newUser)
}

func admin(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "login.gohtml", nil)
}

func login(ctx echo.Context) error {
	render := func(code int, err error) error {
		return ctx.Render(code, "login.gohtml", map[string]interface{}{
			"error": err,
		})
	}

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	conn, err := db.Acquire()
	if err != nil {
		return render(http.StatusInternalServerError, err)
	}
	defer conn.Release()

	users := repo.NewUsers(conn)

	var dbPassword []byte
	dbPassword, err = users.Password(username)
	if err != nil {
		return render(http.StatusInternalServerError, err)
	}
	if dbPassword == nil {
		return render(http.StatusUnauthorized, errors.New("no such user"))
	}

	err = bcrypt.CompareHashAndPassword(dbPassword, []byte(password))
	if err != nil {
		return render(http.StatusUnauthorized, err)
	}

	return ctx.Redirect(http.StatusFound, "/")
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

	passwordHashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	var conn *pgxpool.Conn
	conn, err = db.Acquire()
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
