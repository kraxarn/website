package group

import (
	"fmt"
	"github.com/kraxarn/website/api/teamspeak"
	"github.com/kraxarn/website/helper"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func RegisterServers(app *echo.Echo) {
	group := app.Group("/servers")

	group.GET("", servers)

	group.GET("/teamspeak/status", teamSpeakStatus)
	group.GET("/teamspeak/clients", teamSpeakClients)
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

	err = teamspeak.StatusError(resp)
	if err != nil {
		return err
	}

	str := fmt.Sprintf("%s/%s",
		resp.Body.VirtualServersTotalClientsOnline,
		resp.Body.VirtualServersTotalMaxClients,
	)
	return ctx.String(http.StatusOK, str)
}

func teamSpeakClients(ctx echo.Context) error {
	api, err := teamspeak.NewApi()
	if err != nil {
		return err
	}

	var resp teamspeak.ApiResponse[[]teamspeak.Client]
	resp, err = api.ClientList()
	if err != nil {
		return err
	}

	err = teamspeak.StatusError(resp)
	if err != nil {
		return err
	}

	var builder strings.Builder
	for _, client := range resp.Body {
		builder.WriteString(fmt.Sprintf("%s\n", client.ClientNickname))
	}
	return ctx.String(http.StatusOK, builder.String())
}
