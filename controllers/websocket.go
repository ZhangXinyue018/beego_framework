package controllers

import (
	"github.com/gorilla/websocket"
	"fmt"
	"time"
	"log"
	"bytes"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

type WebSocketController struct {
	MainController
}

type ConnectionClinet struct {
	Connection *websocket.Conn
	send chan []byte
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (webSocketController *WebSocketController) Get() {
	conn, err := Upgrader.Upgrade(
		webSocketController.Ctx.ResponseWriter,
		webSocketController.Ctx.Request,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	connectionClinet := ConnectionClinet{
		Connection: conn,
	}
	go connectionClinet.writeData()
	go connectionClinet.readData()
}

func (conn *ConnectionClinet) writeData() {
	ticker := time.NewTicker(54 * time.Second)
	defer func() {
		ticker.Stop()
		conn.Connection.Close()
	}()
	for {
		select {
		case message, ok := <-conn.send:
			conn.Connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if !ok {
				// The hub closed the channel.
				conn.Connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := conn.Connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(conn.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-conn.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			conn.Connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
			if err := conn.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (conn *ConnectionClinet) readData() {
	defer func() {
		conn.Connection.Close()
	}()
	conn.Connection.SetReadLimit(512)
	conn.Connection.SetReadDeadline(time.Now().Add(60 * time.Second))
	conn.Connection.SetPongHandler(
		func(string) error {
			conn.Connection.SetReadDeadline(time.Now().Add(60 * time.Second));
			return nil
		},
	)
	for {
		_, message, err := conn.Connection.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		fmt.Println(message)
	}
}
