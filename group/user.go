package group

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kraxarn/website/config"
	"github.com/kraxarn/website/data"
	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/repo"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net"
	"net/http"
	"strconv"
	"time"
)

const (
	authMsgGeneric = "server error"
	authMsgLogin   = "invalid username or password"
	authMsgToken   = "invalid token"
	authMsgFlag    = "invalid user"
)

func RegisterUser(app *echo.Echo) {
	group := app.Group("/user")

	group.GET("/login", loginPage)
	group.POST("/login", login)
	group.POST("/new", newUser)
}

func loginPage(ctx echo.Context) error {
	return ctx.Render(http.StatusOK, "login.gohtml", nil)
}

func login(ctx echo.Context) error {
	render := func(code int, message string, err error) error {
		if err != nil {
			ctx.Logger().Error(err)
		}
		return ctx.Render(code, "login.gohtml", map[string]interface{}{
			"error": message,
		})
	}

	username := ctx.FormValue("username")
	password := ctx.FormValue("password")

	conn, err := db.Acquire()
	if err != nil {
		return render(http.StatusInternalServerError, authMsgGeneric, err)
	}
	defer conn.Release()

	users := repo.NewUsers(conn)

	var dbPassword []byte
	dbPassword, err = users.Password(username)
	if err != nil || dbPassword == nil {
		return render(http.StatusUnauthorized, authMsgLogin, nil)
	}

	err = bcrypt.CompareHashAndPassword(dbPassword, []byte(password))
	if err != nil {
		return render(http.StatusUnauthorized, authMsgLogin, err)
	}

	var userId db.Id
	userId, err = users.Id(username)
	if err != nil {
		return render(http.StatusInternalServerError, authMsgGeneric, err)
	}

	// TODO: Maybe fetch both id and user flags at the same time

	var userFlags data.UserFlags
	userFlags, err = users.Flags(userId)
	if err != nil || (userFlags&data.UserFlagsLogin) == 0 {
		return render(http.StatusForbidden, authMsgFlag, err)
	}

	var token config.Token
	token, err = config.NewToken()
	if err != nil {
		return render(http.StatusInternalServerError, authMsgToken, err)
	}

	now := time.Now().UTC()

	var jwtToken string
	jwtToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": jwt.NewNumericDate(now.Add(time.Hour * 24)),
		"nbf": jwt.NewNumericDate(now),
		"sub": strconv.FormatInt(int64(userId), 10),
		"flg": strconv.FormatUint(uint64(userFlags), 10),
	}).SignedString(token.Key())

	if err != nil {
		return render(http.StatusInternalServerError, authMsgToken, err)
	}

	ctx.SetCookie(&http.Cookie{
		Name:     "session",
		Value:    jwtToken,
		Path:     "/admin",
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

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

	_, err = users.Insert(username, passwordHashed[:], data.UserFlagsNone)
	if err != nil {
		return err
	}

	return ctx.NoContent(http.StatusCreated)
}
