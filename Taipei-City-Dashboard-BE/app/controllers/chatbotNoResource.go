package controllers

import (
	// "TaipeiCityDashboardBE/app/models"
	"fmt"
	"strings"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

type NoResourceLocation struct {
	theType string `json:"theType"`
	lat float64 `json:"lat"`
	lng float64 `json:"lng"`
	message string `json:"message"`

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

	// Todo store


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
	if isSuccess{
		repMsg.Message = "have set as no resource message"
	} else{
		repMsg.Message = "error when store no resource message"
	}
	
	return repMsg
}