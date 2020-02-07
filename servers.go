package main

import (
	"fmt"
	"net"
)

type ServerInfo struct {
	Name, Players, Link string
	Online bool
}

type Ts3 struct {
	conn net.Conn
}

func ConnectTs3() (*Ts3, error) {
	// Create instance
	ts3 := new(Ts3)
	// Try to connect
	conn, err := net.Dial("tcp", "kraxarn.com:10011")
	if err != nil {
		return nil, err
	} else {
		ts3.conn = conn
	}
	// Use virtual server 1 by default
	if _, err = fmt.Fprintf(conn, "use 1\n"); err != nil {
		return nil, err
	}
	return ts3, nil
}

func GetServerInfo() []ServerInfo {
	// Make new splice
	infos := make([]ServerInfo, 0)
	// TS3
	ts3, _ := ConnectTs3()
	infos = append(infos, ServerInfo{
		Name:    "TS3",
		Players: "-/-",
		Link:    "-",
		Online:  ts3 != nil,
	})
	// Return final list
	return infos
}