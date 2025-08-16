package group

import (
	"fmt"
	"github.com/kraxarn/website/api/teamspeak"
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RegisterServers(app *echo.Echo) {
	group := app.Group("/servers")

	group.GET("", servers)

	group.GET("/teamspeak/status", teamSpeakStatus)
}

func servers(ctx echo.Context) error {
	return helper.RenderPage(ctx, "servers", map[string]interface{}{
		"styles":  []string{"servers"},
		"scripts": []string{"servers"},
	})
}

func teamSpeakStatus(ctx echo.Context) error {
	api, err := teamspeak.NewApi()
	if err != nil {
		return err
	}

	var resp teamspeak.ApiResponse[teamspeak.HostInfo]
	resp, err = api.HostInfo()
	if err != nil {
		return err
	}

	if err = teamspeak.StatusError(resp); err != nil {
		return err
	}

	str := fmt.Sprintf("%s/%s",
		resp.Body.VirtualServersTotalClientsOnline,
		resp.Body.VirtualServersTotalMaxClients,
	)
	return ctx.String(http.StatusOK, str)
}
