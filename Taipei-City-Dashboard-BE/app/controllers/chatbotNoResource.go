package controllers

import (
	// "TaipeiCityDashboardBE/app/models"
	"fmt"
	"strings"
	"strconv"
	"time"
)

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

func NoResource(repMsg repMessage) repMessage {
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