package service

import (
	"github.com/gorilla/websocket"
	"beego_framework/domain/socket"
	"time"
	"bytes"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WebSocketService struct {
	EventChannels map[string]socket.EventChannel
	ConnectionMap map[*websocket.Conn]*socket.Client
}

func (service *WebSocketService) HandleChannelEvents() () {
	for {
		for _, eventChannel := range service.EventChannels {
			select {
			case broadcast := <-eventChannel.Broadcast:
				for _, client := range eventChannel.Clients {
					client.Send <- broadcast
				}
			case register := <-eventChannel.Register:
				eventChannel.Clients[register.Connection] = register
			case unRegister := <-eventChannel.UnRegister:
				delete(eventChannel.Clients, unRegister.Connection)
			}
		}
	}
}

func (service *WebSocketService) JoinEvent(client *socket.Client, eventName string) () {
	if _, ok := service.EventChannels[eventName]; !ok {
		service.EventChannels[eventName] = socket.EventChannel{
			Clients:    map[*websocket.Conn]*socket.Client{},
			Broadcast:  make(chan socket.Message, 10),
			Register:   make(chan *socket.Client, 10),
			UnRegister: make(chan *socket.Client, 10),
		}
	}
	service.EventChannels[eventName].Register <- client
}

func (service *WebSocketService) LeaveEvent(client *socket.Client, eventName string) () {
	if _, ok := service.EventChannels[eventName]; !ok {
		return
	}
	service.EventChannels[eventName].UnRegister <- client
}

func (service *WebSocketService) CreateClient(client *socket.Client) () {
	service.ConnectionMap[client.Connection] = client
	service.JoinEvent(client, "broadcast")
	go func() {
		if err := client.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
			return
		}
	}()
	go service.keepReading(client)
	go service.keepWriting(client)
}

func (service *WebSocketService) closeClient(client *socket.Client) () {
	defer func() {
		if x := recover(); x != nil {
			//ignore error
		}
	}()
	for _, eventChannel := range service.EventChannels {
		eventChannel.UnRegister <- client
	}
	close(client.Send)
	delete(service.ConnectionMap, client.Connection)
	client.Connection.Close()
}

func (service *WebSocketService) keepReading(client *socket.Client) () {
	defer func() {
		if x := recover(); x != nil {
			service.closeClient(client)
		}
	}()
	for {
		_, message, err := client.Connection.ReadMessage()
		if err != nil {
			service.closeClient(client)
			return
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		service.generateMessages(string(message))
	}
}

func (service *WebSocketService) keepWriting(client *socket.Client) () {
	defer func() {
		if x := recover(); x != nil {
			service.closeClient(client)
		}
	}()
	for {
		select {
		case message := <-client.Send:
			err := client.Connection.WriteJSON(message)
			if err != nil {
				service.closeClient(client)
				return
			}
		}
	}
}

func (service *WebSocketService) generateMessages(msgContent string) () {
	msg := socket.Message{Message: time.Now().Format("2006-01-02 15:04:05") + ": " + msgContent}
	if value, ok := service.EventChannels["broadcast"]; ok {
		value.Broadcast <- msg
	}
}
