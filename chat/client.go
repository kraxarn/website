package chat

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

type Client struct {
	// Hub client belongs to
	hub *Hub
	// WebSocket connection
	connection *websocket.Conn
	// Buffer
	send chan []byte
}

func NewClient(hub *Hub, connection *websocket.Conn) *Client {
	client := new(Client)

	client.hub = hub
	client.connection = connection
	client.send = make(chan []byte, 256)

	client.hub.register <- client
	return client
}

func (client *Client) Read() {
	defer func() {
		client.hub.unregister <- client
		if err := client.connection.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	client.connection.SetReadLimit(maxMessageSize)
	_ = client.connection.SetReadDeadline(time.Now().Add(pongWait))
	client.connection.SetPongHandler(func(appData string) error {
		return client.connection.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, msg, err := client.connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println("error:", err)
			}
			break
		}
		msg = bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
		client.hub.broadcast <- msg
	}
}

func (client *Client) Write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if err := client.connection.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	for {
		select {
		case msg, ok := <-client.send:
			_ = client.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// Hub closed
				_ = client.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			writer, err := client.connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = writer.Write(msg)

			// Add queue
			length := len(client.send)
			for i := 0; i < length; i++ {
				_, _ = writer.Write(newline)
				_, _ = writer.Write(<-client.send)
			}

			if err := writer.Close(); err != nil {
				fmt.Println(err)
				return
			}

		case <-ticker.C:
			_ = client.connection.SetWriteDeadline(time.Now().Add(writeWait))
			if err := client.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
