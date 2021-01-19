package chat

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var webSocketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		fmt.Printf("check origin: %s\n", request.Body)
		return true
	},
	HandshakeTimeout: time.Second * 5,
}
