package controllers

import (
	"fmt"
	"strings"
	"time"
)

// ChatbotDistribute is a function that distributes chatbot messages to the correct chatbot
func ChatbotDistribute(repMsg repMessage) repMessage {
	fmt.Println("ChatbotDistribute")
	fmt.Println(repMsg)
	time.Now()
	newRepMsg := repMsg
	if strings.HasPrefix(repMsg.Message, "!tail ") {

		newRepMsg = AddTail(repMsg)
		fmt.Println("===========add tail===========")
		fmt.Println(repMsg)
	} else if strings.HasPrefix(repMsg.Message, "!a ") || strings.HasPrefix(repMsg.Message, "!w") {

		// is an announcement or a wish
		// start with !a or !w

		// repMsg = a function that returns the repMessage
		fmt.Println("===========store db===========")
		newRepMsg = StoreDB(repMsg)

	} else if strings.HasPrefix(repMsg.Message, "!g "){
		fmt.Println("===========gemini===========")
		newRepMsg = AskGemini(repMsg)
		
	} else if strings.HasPrefix(repMsg.Message, "!no "){
		fmt.Println("===========no===========")
		newRepMsg = NoResourceC(repMsg)
	} else{
		fmt.Println("===========else===========")
		newRepMsg = repMsg
		newRepMsg.Message = "Find no command, please try again."
		
	}

	newRepMsg.Username = "DashBoardbot"


	return newRepMsg

}
