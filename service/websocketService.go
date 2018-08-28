package service

import (
	"github.com/gorilla/websocket"
	"beego_framework/domain/socket"
	"fmt"
	"time"
	"bytes"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
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
						go func() {
							err := connection.WriteJSON(broadcast)
							if err != nil {
								service.closeConn(connection)
							}
						}()
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
	go func() {
		if err := conn.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
	}()
	go service.keepReading(conn)
}

func (service *WebSocketService) closeConn(conn *websocket.Conn) () {
	for _, eventChannel := range service.EventChannels {
		eventChannel.UnRegister <- conn
	}
	conn.Close()
}

func (service *WebSocketService) keepReading(conn *websocket.Conn) () {
	defer func() {
		service.closeConn(conn)
	}()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			service.closeConn(conn)
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		fmt.Println("Received message from client: " + string(message))
		service.generateMessages(string(message))
	}
}

func (service *WebSocketService) generateMessages(msgContent string) () {
	msg := socket.Message{Message: time.Now().Format("2006-01-02 15:04:05") + ": " + msgContent}
	if value, ok := service.EventChannels["broadcast"]; ok {
		value.Broadcast <- msg
	}
}
