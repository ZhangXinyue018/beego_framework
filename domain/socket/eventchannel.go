package socket

import (
	"github.com/gorilla/websocket"
)

type EventChannel struct {
	Clients    map[*websocket.Conn]*Client
	Broadcast  chan Message
	Register   chan *Client
	UnRegister chan *Client
}

type Client struct {
	Connection *websocket.Conn
	Send       chan Message
}
