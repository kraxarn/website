package chat

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func Route(router *gin.Engine) {
	var hubs map[string]Hub

	router.GET("/chat/:id", func(context *gin.Context) {
		id := context.Param("id")
		hub, found := hubs[id]
		if !found {
			fmt.Printf("creating hub: \"%s\"\n", id)
			hub = NewHub()
			hubs[id] = hub
		}

		hub.Serve(context.Writer, context.Request)
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
