package controllers

import (
	"github.com/gorilla/websocket"
	"fmt"
		"net/http"
	"beego_framework/bean"
		)

type WebSocketController struct {
	MainController
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
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

