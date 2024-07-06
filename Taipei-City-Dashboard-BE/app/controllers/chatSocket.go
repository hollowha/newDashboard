package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan revMessage)

// client to server message
type revMessage struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}



// server to client message

// UserDisplay
// -1: no display
// 0: all
// UserID (1~): display to specific user

// DashboardDisplay
// -1: no display
// 0: all
// DashboardID (1~): display to specific dashboard

// message type
// announcement // into db bg-color red
// wish // into db bg-color blue
// message // do Not store in db color white

// component
// the list

type repMessage struct {
	UserDisplay int `json:"userDisplay"`
	DashboardDisplay int `json:"dashboardDisplay"`
	MessageType string `json:"type"`
	Username string `json:"username"`
	Message  string `json:"message"`
	Component []int `json:"component"`

}

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("Error upgrading to websocket: %v", err)
	}
	defer ws.Close()
	clients[ws] = true
	for {
		var msg revMessage
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error reading json: %v", err)
			delete(clients, ws)
			break
		}


		broadcast <- msg

	}
}

func SendMessage(msg repMessage) {
	// send the message to all clients
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


func HandleMessages() {
	for {
		msg := <-broadcast

		// Convert revMessage to repMessage before sending to the broadcast channel
		repMsg := repMessage{
			UserDisplay:      0,
			DashboardDisplay: 0,
			MessageType:      "message",
			Username:         msg.Username,
			Message:          msg.Message,
			Component:        []int{5, 10, 43, 69, 7},
		}


		// revieve the message
		
		if strings.HasPrefix(msg.Message, "!") {
			// is a command
			// first send the command itself to the client
			SendMessage(repMsg)
	
			// then process the command
			// is an announcement or a wish
			// start with !a or !w
	
			// repMsg = a function that returns the repMessage
			// SendMessage(repMsg)
		} else {
			// is a message
			SendMessage(repMsg)
		}
		fmt.Println(msg)

		fmt.Println(repMsg)
		fmt.Println("========")

		

		
	}
}
