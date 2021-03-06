package bean

import (
	"beego_framework/domain/socket"
	"beego_framework/service"
	serviceimpl "beego_framework/service/impl"
	"github.com/gorilla/websocket"
)

func InitService() {
	ExchangerServiceBean = &serviceimpl.ExchangerService{
		ExchangerRpc: ExchangerRpcBean,
	}
	WebSocketServiceBean = &service.WebSocketService{
		EventChannels: map[string]*socket.EventChannel{},
		ConnectionMap: map[*websocket.Conn]*socket.Client{},
	}
	TestServiceBean = &serviceimpl.TestService{
		Repository: MysqlTempRepoBean,
	}
}
