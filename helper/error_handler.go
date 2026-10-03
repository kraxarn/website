package helper

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/kraxarn/website/config"
	"github.com/labstack/echo/v5"
)

func HandleError(ctx *echo.Context, err error) {
	if resp, uwErr := echo.UnwrapResponse(ctx.Response()); uwErr == nil {
		if resp.Committed {
			return
		}
	}

	var code int
	var httpErr *echo.HTTPError

	if errors.As(err, &httpErr) {
		code = httpErr.Code
	} else {
		code = http.StatusInternalServerError
		if err != nil {
			ctx.Logger().Error("request failed", "error", err)
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
			ctx.Logger().Error("string failed", "error", err)
		}

		return
	}

	err = ctx.Render(code, "error.gohtml", map[string]interface{}{
		"StatusCode": code,
		"StatusText": http.StatusText(code),
	})
	if err != nil {
		ctx.Logger().Error("render failed", "error", err)
	}
}
