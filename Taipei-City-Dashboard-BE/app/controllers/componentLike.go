package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func LikeComponentByID(c *gin.Context) {
	// Get the user ID from the context
	userID := c.GetInt("accountID")
	// Get the component ID from the form
	componentIDstr := c.PostForm("componentid")
	fmt.Println(userID , " Like commponent ", componentIDstr, " <3<3<3")

	// turn string to int
	componentid, err := strconv.Atoi(componentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid component ID"})
		return
	}
	response := fmt.Sprintf("User %d liked component %d", userID, componentid)

	// Database modify

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func GetLikeComponentByID(c *gin.Context) {
	// Get the component ID from the form
	componentIDstr := c.Param("componentid")
	fmt.Println("Get like commponent id=", componentIDstr, " <3<3<3")

	// turn string to int
	// unmark this if you want to use the componentid: int
	/*
	componentid, err := strconv.Atoi(componentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid component ID"})
		return
	}
	*/

	// Database get
	// Get the like count from the database
	
	// temp value
	likeCount := 5

	

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": likeCount})
}