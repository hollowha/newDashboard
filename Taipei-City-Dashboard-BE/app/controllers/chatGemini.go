package controllers

import (
	"fmt"
)


func AskGemini( repMsg repMessage) repMessage {
	fmt.Println("AskGemini")
	fmt.Println(repMsg)

	return repMsg


}
