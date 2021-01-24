package chat

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/config"
	"github.com/kraxarn/website/user"
	"strings"
)

type Hub struct {
	// Clients
	clients map[*Client]bool
	// Inbound message
	broadcast chan []byte
	// Register requests
	register chan *Client
	// Unregister requests
	unregister chan *Client
	// Token for user management
	token *config.Token
}

func NewHub(token *config.Token) Hub {
	hub := Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		token:      token,
	}
	go hub.Run()
	return hub
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register:
			hub.clients[client] = true

		case client := <-hub.unregister:
			if _, ok := hub.clients[client]; ok {
				delete(hub.clients, client)
				close(client.send)
			}

		case message := <-hub.broadcast:
			for client := range hub.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(hub.clients, client)
				}
			}
		}
	}
}

func (hub *Hub) Serve(ctx *gin.Context) {
	connection, err := webSocketUpgrade.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	usr, err := user.NewUserFromCookie(ctx, hub.token)
	if err != nil {
		return
	}

	client := NewClient(hub, connection, usr)
	go client.Write()
	go client.Read()
}
