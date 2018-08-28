package bean

import (
	"beego_framework/service"
	"beego_framework/domain/socket"
	"github.com/gorilla/websocket"
)

var (
	ExchangerServiceBean *service.ExchangerService
	WebSocketServiceBean *service.WebSocketService
)

func init() {
	ExchangerServiceBean = &service.ExchangerService{
		ExchangerRpc: ExchangerRpcBean,
	}
	WebSocketServiceBean = &service.WebSocketService{
		EventChannels: map[string]socket.EventChannel{},
		ConnectionMap: map[*websocket.Conn]bool{},
	}
}
