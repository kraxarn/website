package main

import (
	"fmt"
	"github.com/kraxarn/website/config"
	"github.com/kraxarn/website/db"
	"github.com/kraxarn/website/group"
	"github.com/kraxarn/website/helper"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	app := echo.New()

	if err := initMiddleware(app); err != nil {
		app.Logger.Fatal(err)
	}

	initGroups(app)

	renderer, err := helper.NewTemplateRenderer()
	if err != nil {
		app.Logger.Fatal(err)
	}

	app.Renderer = renderer
	app.HTTPErrorHandler = helper.HandleError

	if err = db.Connect(); err != nil {
		app.Logger.Fatal(err)
	}
	defer db.Close()

	if err := app.Start("127.0.0.1:5000"); err != nil {
		app.Logger.Fatal(err)
	}
}

func initMiddleware(app *echo.Echo) error {
	app.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus: true,
		LogURI:    true,
		LogValuesFunc: func(ctx echo.Context, val middleware.RequestLoggerValues) error {
			var writer io.Writer
			if val.Error != nil {
				writer = os.Stderr
			} else {
				writer = os.Stdout
			}

			_, err := fmt.Fprintf(writer,
				"%-20s %-4d %-6s %-16s %-8s %s\n",
				val.StartTime.Format(time.DateTime), // 9999-12-31 23:59:59 (19)
				val.Status,                          // 999 (3)
				val.Latency.String(),                // 999ms (5)
				ctx.RealIP(),                        // 255.255.255.255 (15)
				ctx.Request().Method,                // OPTIONS (7)
				val.URI,                             // / (...)
			)

			return err
		},
	}))

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "static",
		Browse:     false,
		IgnoreBase: true,
	}))

	// Default rate limiter
	app.Use(middleware.RateLimiterWithConfig(
		middleware.RateLimiterConfig{
			Skipper: func(ctx echo.Context) bool {
				return strings.HasPrefix(ctx.Path(), "/admin")
			},
			Store: middleware.NewRateLimiterMemoryStore(
				rate.Limit(10),
			),
		},
	))

	// Admin rate limiter
	app.Use(middleware.RateLimiterWithConfig(
		middleware.RateLimiterConfig{
			Skipper: func(ctx echo.Context) bool {
				return !strings.HasPrefix(ctx.Path(), "/admin")
			},
			Store: middleware.NewRateLimiterMemoryStore(
				rate.Limit(1),
			),
		},
	))

	token, err := config.NewToken()
	if err != nil {
		return err
	}
	app.Use(echojwt.JWT(token.Key()))

	app.Use(middleware.Recover())

	return nil
}

func initGroups(app *echo.Echo) {
	group.RegisterIndex(app)
	group.RegisterAdmin(app)
}
