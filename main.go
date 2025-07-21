package main

import (
	"fmt"
	"github.com/kraxarn/website/group"
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
	"io"
	"os"
	"time"
)

func main() {
	app := echo.New()

	initMiddleware(app)
	initGroups(app)

	renderer, err := helper.NewTemplateRenderer()
	if err != nil {
		app.Logger.Fatal(err)
	}

	app.Renderer = renderer
	app.HTTPErrorHandler = helper.HandleError

	if err := app.Start("127.0.0.1:5000"); err != nil {
		app.Logger.Fatal(err)
	}
}

func initMiddleware(app *echo.Echo) {
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

	app.Use(middleware.RateLimiter(
		middleware.NewRateLimiterMemoryStore(
			rate.Limit(10),
		),
	))

	app.Use(middleware.Recover())
}

func initGroups(app *echo.Echo) {
	group.RegisterIndex(app)
}
