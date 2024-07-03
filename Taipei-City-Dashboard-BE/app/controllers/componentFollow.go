package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowComponentByID(c *gin.Context) {
	// Get the user ID from the context
	userID := c.GetInt("accountID")
	// Get the component ID from the form
	componentIDstr := c.PostForm("componentid")
	fmt.Println(userID , " want to Follow commponent ", componentIDstr, " <3<3<3")

	// turn string to int
	componentid, err := strconv.Atoi(componentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid component ID"})
		return
	}
	response := fmt.Sprintf("User %d want to Follow component %d", userID, componentid)

	// Database modify

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func GetFollowComponentListByUserID(c *gin.Context) {
	// Get the user ID from the context
	userID := c.GetInt("accountID")
	fmt.Println("Get Follow commponent list by user id=", userID, " <3<3<3")

	// Database get
	// Get the follow list from the database
	
	// temp value
	followList := []int{1, 2, 3, 5, 8, 13}

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": followList})
}