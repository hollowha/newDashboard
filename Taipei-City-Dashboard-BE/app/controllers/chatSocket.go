package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
	"TaipeiCityDashboardBE/app/models"
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
	DashboardDisplay string `json:"dashboardDisplay"`
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
	UserDisplay       string    `json:"userDisplay" gorm:"column:user_display"`
	DashboardDisplay  string       `json:"dashboardDisplay" gorm:"column:dashboard_display"`
	MessageType       string    `json:"type" gorm:"column:message_type"`
	Username          string    `json:"username" gorm:"column:user_name"`
	Message           string    `json:"message" gorm:"column:message"`
	// Component         []int     `json:"component" gorm:"column:component"`
	Timestamp         time.Time `json:"timestamp" gorm:"column:upload_time"`
}


func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatalf("Error upgrading to websocket: %v", err)
	}
	defer ws.Close()
	clients[ws] = true

	var initialMessages []repMessage
	table := models.DBManager.Table("rep_message")
	if err := table.Find(&initialMessages).Error; err != nil {
        log.Printf("Error retrieving messages from database: %v", err)
        return
    }

	fmt.Println("start send announce message")
	for _, msg := range initialMessages {
		if msg.MessageType == "announcement" {
			fmt.Printf("send announce message: %v\n", msg)
			err = ws.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending initial message: %v", err)
				ws.Close()
				delete(clients, ws)
				return
			}

		}
		
    }
	fmt.Println("start send wish message done")
	for _, msg := range initialMessages {
		
		if msg.MessageType == "wish" {
			fmt.Printf("send wish message: %v\n", msg)
			err = ws.WriteJSON(msg)
			if err != nil {
				log.Printf("Error sending initial message: %v", err)
				ws.Close()
				delete(clients, ws)
				return
			}
		}
    }
	fmt.Println("send init message done")

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
			UserDisplay:      "0",
			DashboardDisplay: msg.DashboardDisplay,
			MessageType:      "message",
			Username:         msg.Username,
			Message:          strings.TrimSpace(msg.Message),
			// Component:        []int{5, 10, 43, 69, 7},
			Timestamp:        time.Now(),
		}
		// revieve the message
		
		if strings.HasPrefix(msg.Message, "!") {
			// is a command
			// first send the command itself to the client
			SendMessage(repMsg)
			newRepMsg := ChatbotDistribute(repMsg)

			SendMessage(newRepMsg)
		} else {
			// is a message
			SendMessage(repMsg)
		}
		fmt.Println(msg)
		fmt.Println(repMsg)
		fmt.Println("========")
	}
}
