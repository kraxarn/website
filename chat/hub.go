package chat

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kraxarn/website/config"
	"github.com/kraxarn/website/user"
	"strings"
)

type ClientMessage struct {
	Message []byte
	Sender  *Client
}

type Hub struct {
	// Clients
	clients map[*Client]bool
	// Inbound message
	broadcast chan ClientMessage
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
		broadcast:  make(chan ClientMessage),
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
			msg, err := json.Marshal(hub.getBroadcastMessage(message))
			if err != nil {
				break
			}

			for client := range hub.clients {
				select {
				case client.send <- msg:
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

func (hub *Hub) getBroadcastMessage(message ClientMessage) Message {
	sender := message.Sender.user
	msg := strings.Split(string(message.Message), " ")

	switch msg[0] {
	case "/video":
		if len(msg) >= 2 {
			if msg, err := NewVideoMessage(sender, msg[1]); err == nil {
				return msg
			} else {
				fmt.Println(err)
			}
		}
	}

	return NewMessage(sender, string(message.Message))
}
