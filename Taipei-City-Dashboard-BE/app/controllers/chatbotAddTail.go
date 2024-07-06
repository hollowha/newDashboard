package controllers

import (
	"fmt"
	"strings"
)


func AddTail(repMsg repMessage) repMessage {
	fmt.Println("AddTail")
	
	part := strings.SplitN(repMsg.Message, " ", 2)
	
	repMsg.Message = part[1] + " tail~~~~~~~~~~~"
	return repMsg

}