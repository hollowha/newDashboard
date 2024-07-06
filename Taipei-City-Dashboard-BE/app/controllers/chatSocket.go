package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("Error upgrading to websocket: %v", err)
	}
	defer ws.Close()
	clients[ws] = true
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading json: %v", err)
			delete(clients, ws)
			break
		}
		broadcast <- msg
	}
}


func HandleMessages() {
	for {
		msg := <-broadcast

		// message will come into here in mesg
		// check the form of the messege decide whether it is the command or not
		fmt.Println(msg)

		// sent the message back to every client
		for client := range clients {
			// send
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Error writing json: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
