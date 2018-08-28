package service

import (
	"github.com/gorilla/websocket"
	"beego_framework/domain/socket"
	"fmt"
)

type WebSocketService struct {
	EventChannels map[string]socket.EventChannel
	ConnectionMap map[*websocket.Conn]bool
}

func (service *WebSocketService) HandleChannelEvents() () {
	for {
		for _, eventChannel := range service.EventChannels {
			select {
			case broadcast := <-eventChannel.Broadcast:
				for connection, value := range eventChannel.Clients {
					if value {
						err := connection.WriteJSON(broadcast)
						if err != nil {
							fmt.Println(err)
							service.CloseConn(connection)
						}
					}
				}
			case register := <-eventChannel.Register:
				eventChannel.Clients[register] = true
			case unRegister := <-eventChannel.UnRegister:
				delete(eventChannel.Clients, unRegister)
			}
		}
	}
}

func (service *WebSocketService) JoinEvent(conn *websocket.Conn, eventName string) () {
	if _, ok := service.EventChannels[eventName]; !ok {
		service.EventChannels[eventName] = socket.EventChannel{
			Clients:    map[*websocket.Conn]bool{},
			Broadcast:  make(chan socket.Message, 10),
			Register:   make(chan *websocket.Conn, 10),
			UnRegister: make(chan *websocket.Conn, 10),
		}
	}
	service.EventChannels[eventName].Register <- conn
}

func (service *WebSocketService) CreateConn(conn *websocket.Conn) () {
	service.ConnectionMap[conn] = true
	service.JoinEvent(conn, "broadcast")
}

func (service *WebSocketService) CloseConn(conn *websocket.Conn) () {
	for _, eventChannel := range service.EventChannels {
		eventChannel.UnRegister <- conn
	}
	conn.Close()
}
