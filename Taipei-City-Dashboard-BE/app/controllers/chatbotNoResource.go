package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type NoResourceLocation struct {
	message string    `json:"message" gorm:"column:content"`
	theType string    `json:"theType" gorm:"column:type"`
	theTime time.Time `json:"theTime" gorm:"column:time"`
	lat     float64   `json:"lat" gorm:"column:lat"`
	lng     float64   `json:"lng" gorm:"column:lng"`
}

type NoResourceLocationWithoutTime struct {
	TheType string  `json:"theType" gorm:"type"`
	Lat     float64 `json:"lat" gorm:"lat"`
	Lng     float64 `json:"lng" gorm:"lng"`
	Message string  `json:"message" gorm:"content"`
}


func StoreNoResourceLocation(theType string, lat float64, lng float64, message string) bool {
	fmt.Println("-------StoreNoResourceLocation--------")
	happenTime := time.Now()
	// theType: elec, water, gas, road, other
	// lat: double
	// lng: double
	// message: string may be any thing
	// happenTime: just the time

	fmt.Println(happenTime)
	query := `
		INSERT INTO public.report (content, type, "time", lng, lat)
		VALUES (?, ?, ?, ?, ?)
	`
	result := models.DBDashboard.Exec(query, message, theType, happenTime, lng, lat)

	if result.Error != nil {
		fmt.Println("Error inserting data:", result.Error)
		return false
	}

	return true
}

// for RESTFUL API
func NoResourceR(c *gin.Context) {
	fmt.Println("-------NoResource RESTFUL--------")
	// input as a json
	// {
	// 	"theType": "WaterSource",
	// 	"lat": 35.6895,
	// 	"lng": 139.6917,
	// 	"message": "No water source available here."
	// }

	var noResourceLocationWithoutTime NoResourceLocationWithoutTime
	// bind the json to the struct
	c.BindJSON(&noResourceLocationWithoutTime)
	fmt.Println(noResourceLocationWithoutTime)
	StoreNoResourceLocation(noResourceLocationWithoutTime.TheType, noResourceLocationWithoutTime.Lat, noResourceLocationWithoutTime.Lng, noResourceLocationWithoutTime.Message)
	return
}

// for chatbot
func NoResourceC(repMsg repMessage) repMessage {
	fmt.Println("-------NoResource--------")
	// form: !no elec 21.1234 121.1234 here does not have eletricity
	part := strings.SplitN(repMsg.Message, " ", 5)
	theType := part[1]

	lat, err := strconv.ParseFloat(part[2], 64)
	if err != nil {
		fmt.Println("转换错误:", err)
		repMsg.Message = "error when store no resource message"
		return repMsg
	}
	lng, err := strconv.ParseFloat(part[3], 64)
	if err != nil {
		fmt.Println("转换错误:", err)
		repMsg.Message = "error when store no resource message"
		return repMsg
	}
	message := part[4]
	fmt.Println(message)
	isSuccess := StoreNoResourceLocation(theType, lat, lng, message)
	if isSuccess {
		repMsg.Message = "have set as no resource message"
	} else {
		repMsg.Message = "error when store no resource message"
	}

	return repMsg
}
