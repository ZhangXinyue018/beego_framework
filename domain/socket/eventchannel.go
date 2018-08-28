package socket

import (
	"github.com/gorilla/websocket"
)

type EventChannel struct {
	Clients    map[*websocket.Conn]bool
	Broadcast  chan Message
	Register   chan *websocket.Conn
	UnRegister chan *websocket.Conn
}
