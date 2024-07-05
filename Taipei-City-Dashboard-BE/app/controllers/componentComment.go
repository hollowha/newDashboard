package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// a comment struct with comment id, user id, component id, comment content, comment time
type Comment struct {
	CommentID    int    `json:"commentid"`
	UserID       int    `json:"userid"`
	UserName     string `json:"username"`
	ComponentID  int    `json:"componentid"`
	Comment      string `json:"comment"`
	CommentTime  string `json:"commenttime"`
}

func CommentComponentByID(c *gin.Context) {
	// Get the user ID from the context
	userID := c.GetInt("accountID")
	// Get the component ID from the form
	componentIDstr := c.PostForm("componentid")
	comment := c.PostForm("comment")
	fmt.Println(userID , " comment ", componentIDstr, " with \" ", comment," \" <3<3<3")

	// turn string to int
	componentid, err := strconv.Atoi(componentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid component ID"})
		return
	}
	
	// 前端只會傳componentid, comment, 後端自動獲得 userid, commenttime, 資料庫產生 commentid
	// UserName 再討論，或許也可以從資料庫拿

	// get time of now
	commentTime := time.Now().Add(8 * time.Hour).Format("2006-01-02 15:04:05") // UTC+8
	response := fmt.Sprintf("User %d comment to component %d with \" %s \" at %s", userID, componentid, comment, commentTime)

	// Database modify

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

func GetCommentComponentByID(c *gin.Context) {
	componentIDstr := c.Param("componentid")
	fmt.Println("Get Comment from commponent ", componentIDstr, " <3<3<3")

	// Database get
	// Get the follow list from the database
	
	// user name 在想一下怎麼取

	// temp value
	commentList := []Comment{
		{CommentID: 1, UserID: 1, UserName: "user1", ComponentID: 1, Comment: "comment1", CommentTime: "2021-06-01 12:00:00"},
		{CommentID: 2, UserID: 2, UserName: "user2", ComponentID: 1, Comment: "comment2", CommentTime: "2021-06-02 12:00:00"},
		{CommentID: 3, UserID: 3, UserName: "user3", ComponentID: 1, Comment: "comment3", CommentTime: "2021-06-03 12:00:00"},
		{CommentID: 4, UserID: 4, UserName: "user4", ComponentID: 1, Comment: "comment4", CommentTime: "2021-06-04 12:00:00"},
		{CommentID: 5, UserID: 5, UserName: "user5", ComponentID: 1, Comment: "comment5", CommentTime: "2021-06-05 12:00:08"},
	}

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"status": "success", "data": commentList})
}