package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"TaipeiCityDashboardBE/app/models"

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
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid component ID: " + componentIDstr})
        return
    }

    // Database modify
    likeOrNot, er := UpdateLikes(userID, componentid)
    if er != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Component ID: "+componentIDstr+" not found"})
        return
    }
    // form the response
    var response string
    if likeOrNot { // like
        response = fmt.Sprintf("User %d liked component %d", userID, componentid)
    } else { // dislike
        response = fmt.Sprintf("User %d withdrew like from component %d", userID, componentid)
    }
    // Return the success message
    c.JSON(http.StatusOK, gin.H{"status": "success", "data": response})
}

// method to modify table
func UpdateLikes(userID, componentID int) (bool, error) {
    var actionPerformed bool
    table := models.DBManager.Table("users_like") // table: users_like
    err := table.Raw("SELECT update_user_like(?, ?)", userID, componentID).Scan(&actionPerformed).Error
    // result := table.Exec("SELECT update_user_like(?, ?)", userID, componentID)
    return actionPerformed, err
}