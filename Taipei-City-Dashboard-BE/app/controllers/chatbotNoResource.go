package controllers

import (
	// "TaipeiCityDashboardBE/app/models"
	"TaipeiCityDashboardBE/app/models"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type NoResourceLocation struct {
	theType string    `json:"theType" gorm:"type"`
	lat     float64   `json:"lat" gorm:"lat"`
	lng     float64   `json:"lng" gorm:"lng"`
	message string    `json:"message" gorm:"content"`
	theTime time.Time `json:"theTime" gorm:"time"`
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

	// 创建一个 NoResourceLocation 实例并填充数据
	report := NoResourceLocation{
		theType: theType,
		lat:     lat,
		lng:     lng,
		message: message,
		theTime: happenTime,
	}

	// 使用 GORM 将数据插入到数据库中
	result := models.DBDashboard.Table("report").Create(&report)
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
	// form: !no elec 21.1234 121.1234 here does not have eletricity

	var noResourceLocation NoResourceLocation

	if err := c.ShouldBindJSON(&noResourceLocation); err != nil {
		c.JSON(200, gin.H{
			"message": "error when store no resource message",
		})
	}
	fmt.Println(noResourceLocation)
	StoreNoResourceLocation(noResourceLocation.theType, noResourceLocation.lat, noResourceLocation.lng, noResourceLocation.message)
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
