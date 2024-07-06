package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"fmt"
	"strings"
)

func StoreDB(repMsg repMessage) repMessage {
	fmt.Println("-------StoreDB--------")
	if strings.HasPrefix(repMsg.Message, "!a") {
		repMsg.MessageType = "announcement"
	}

	if strings.HasPrefix(repMsg.Message, "!w") {
		repMsg.MessageType = "wish"
	}

	part := strings.SplitN(repMsg.Message, " ", 2)
	
	repMsg.Message = part[1]

	if repMsg.MessageType == "announcement" || repMsg.MessageType == "wish"{
		// TODO store db
		table := models.DBManager.Table("rep_message")
		// store
		err := table.Create(&repMsg).Error
		if err != nil {
			fmt.Println("Error when store rep_message into database")
		}

	}
	repMsg.Message = "have set as " + repMsg.MessageType + " message: " + repMsg.Message
	return repMsg
}
