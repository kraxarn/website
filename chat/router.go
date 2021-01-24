package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/kraxarn/website/common"
	"net/http"
)

func Route(router *gin.Engine) {
	hubs := make(map[string]Hub)

	router.GET("/chat/hub/:id", func(context *gin.Context) {
		id := context.Param("id")
		hub, found := hubs[id]
		if !found {
			fmt.Printf("creating hub: \"%s\"\n", id)
			hub = NewHub()
			hubs[id] = hub
		}

		hub.Serve(context.Writer, context.Request)
	})

	router.GET("/chat/info/:id", func(context *gin.Context) {
		hub, found := hubs[context.Param("id")]
		if !found {
			context.JSON(http.StatusOK, common.NewError(fmt.Errorf("hub not found")))
			return
		}

		context.JSON(http.StatusOK, map[string]interface{}{
			"client_count": len(hub.clients),
		})
	})

	router.GET("/chat/exists/:id", func(context *gin.Context) {
		_, found := hubs[context.Param("id")]
		context.JSON(http.StatusOK, map[string]interface{}{
			"exists": found,
		})
	})
}

func handleWebsocket(context *gin.Context) {
	connection, err := webSocketUpgrade.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		msgType, data, err := connection.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		if msgType != websocket.TextMessage {
			continue
		}
		fmt.Printf("message: %s\n", data)
	}
}
