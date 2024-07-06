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

	if repMsg.MessageType == "announcement" {
		// TODO store db
		table := models.DBManager.Table("rep_message")
		// store
		err := table.Create(&repMsg).Error
		if err != nil {
			fmt.Println("Error when store rep_message into database")
		}

	}

	return repMsg
}
