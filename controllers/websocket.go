package controllers

import (
	"github.com/gorilla/websocket"
	"fmt"
	"time"
	"net/http"
	"beego_framework/bean"
	"beego_framework/domain/socket"
)

type WebSocketController struct {
	MainController
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func init() {
	go func() {
		for{
			generateMessages()
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 3)
			fmt.Println(bean.WebSocketServiceBean.EventChannels)
		}
	}()
	go bean.WebSocketServiceBean.HandleChannelEvents()
}

// @router / [get]
func (webSocketController *WebSocketController) Get() {
	conn, err := upgrader.Upgrade(
		webSocketController.Ctx.ResponseWriter,
		webSocketController.Ctx.Request,
		nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	bean.WebSocketServiceBean.CreateConn(conn)
}

func generateMessages() {
	time.Sleep(time.Second * 3)
	msg := socket.Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
	if value, ok := bean.WebSocketServiceBean.EventChannels["broadcast"]; ok {
		value.Broadcast <- msg
	}
}
