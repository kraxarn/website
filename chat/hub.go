package chat

import (
	"fmt"
	"net/http"
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
}

func NewHub() Hub {
	return Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
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

func (hub *Hub) Serve(writer http.ResponseWriter, request *http.Request) {
	connection, err := webSocketUpgrade.Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := NewClient(hub, connection)
	go client.Write()
	go client.Read()
}
