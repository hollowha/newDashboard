package controllers

import (
	"fmt"
	"strings"
)


func AddTail(repMsg repMessage) repMessage {
	fmt.Println("AddTail")
	
	part := strings.Split(repMsg.Message, " ")
	
	repMsg.Message = part[1] + " tail~~~~~~~~~~~"
	return repMsg

}