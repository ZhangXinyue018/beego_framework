package service

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"beego_framework/domain/socket"
)

type WebSocketService struct {
	ClientChannelMap map[*websocket.Conn]map[string]bool
	EventChannels    map[string]socket.EventChannel
}

func (service *WebSocketService) HandleChannelEvents() () {
	for {
		for eventName, eventChannel := range service.EventChannels {
			select {
			case broadcast := <-eventChannel.Broadcast:
				message, err := json.Marshal(broadcast)
				if err != nil {
					continue
				}
				for connection, value := range eventChannel.Clients {
					if value {
						if connection.WriteMessage(websocket.TextMessage, message) != nil {
							//ignore
						}
					} else {
						eventChannel.UnRegister <- connection
					}
				}
			case register := <-eventChannel.Register:
				eventChannel.Clients[register] = true
				service.ClientChannelMap[register][eventName] = true
			case unRegister := <-eventChannel.UnRegister:
				delete(eventChannel.Clients, unRegister)
				delete(service.ClientChannelMap[unRegister], eventName)
			}
		}
	}
}
