package chat

import "time"

const (
	// Time allowed to write to peer
	writeWait = 10 * time.Second
	// Time to read response from peer
	pongWait = 60 * time.Second
	// Period to send pings, less than pongWait
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size
	maxMessageSize = 512
)
