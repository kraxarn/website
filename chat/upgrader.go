package chat

import (
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var webSocketUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(request *http.Request) bool {
		return true
	},
	HandshakeTimeout: time.Second * 5,
}
