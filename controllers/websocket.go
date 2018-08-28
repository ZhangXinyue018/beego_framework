package controllers

import (
	"github.com/gorilla/websocket"
	"fmt"
	"log"
	"time"
	"net/http"
)

type WebSocketController struct {
	MainController
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Message string `json:"message"`
}

func init() {
	go handleMessages()
	go func() {
		for {
			time.Sleep(time.Second * 3)
			msg := Message{Message: "这是向页面发送的数据 " + time.Now().Format("2006-01-02 15:04:05")}
			broadcast <- msg
		}
	}()
	go func() {
		for {
			time.Sleep(time.Second * 3)
			fmt.Println(clients)
		}

	}()
}

// @router / [get]
func (webSocketController *WebSocketController) Get() {
	fmt.Println("Enter!!!")
	conn, err := upgrader.Upgrade(
		webSocketController.Ctx.ResponseWriter,
		webSocketController.Ctx.Request,
		nil)
	if err != nil {
		fmt.Println("Encounter error!!!")
		fmt.Println(err)
		return
	}

	clients[conn] = true
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("client.WriteJSON error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
