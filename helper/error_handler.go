package helper

import (
	"errors"
	"fmt"
	"github.com/kraxarn/website/config"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func HandleError(err error, ctx echo.Context) {
	if ctx.Response().Committed {
		return
	}

	var code int
	var httpErr *echo.HTTPError

	if errors.As(err, &httpErr) {
		code = httpErr.Code
	} else {
		code = http.StatusInternalServerError
		if err != nil {
			ctx.Logger().Error(err)
		}
	}

	if config.Dev() {
		var builder strings.Builder

		builder.WriteString(fmt.Sprintf("%d: %s", code, http.StatusText(code)))

		if err != nil {
			builder.WriteRune('\n')
			builder.WriteString(err.Error())
		}

		err = ctx.String(code, builder.String())
		if err != nil {
			ctx.Logger().Error(err)
		}

		return
	}

	err = ctx.Render(code, "error.gohtml", map[string]interface{}{
		"StatusCode": code,
		"StatusText": http.StatusText(code),
	})
	if err != nil {
		ctx.Logger().Error(err)
	}
}
