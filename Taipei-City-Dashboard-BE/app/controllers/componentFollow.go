package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FollowDashboardByIndex(c *gin.Context) {
	// Get the user ID from the context
	userID := c.GetInt("accountID")
	// Get the dashboard index from the URL parameter
	dashboardIndex := c.Param("index")
	fmt.Println(userID, " wants to update follow status on Dashboard: ", dashboardIndex)

	// Update user's follow status
	isFollow, err := update_user_follow(userID, dashboardIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Failed to update status of dashboard following"})
		fmt.Println(http.StatusBadRequest, err)
		return
	}

	// Fetch components of followed dashboards
	findComponentsQuery := `
		SELECT DISTINCT unnest(components) AS component
		FROM dashboards
		WHERE index IN (
			SELECT dashboard_index
			FROM user_followed
			WHERE user_id = ?
		);
	`
	var dashboard models.Dashboard
	err = models.DBManager.Raw(findComponentsQuery, userID).Scan(&dashboard.Components).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Failed to get components of followed dashboards"})
		fmt.Println(http.StatusBadRequest, err)
		return
	}

	// Prepare followed dashboard index for update
	var followedDashboardIndex string
	query := `
		SELECT index
		FROM dashboards
		WHERE id IN (
			SELECT dashboard_id
			FROM dashboard_groups
			WHERE group_id = (
				SELECT DISTINCT group_id
				FROM auth_user_group_roles t
				WHERE t.group_id <> 1 AND auth_user_id = ?
			)
		) AND name = '個人追蹤儀表板';
	`
	err = models.DBManager.Raw(query, userID).First(&followedDashboardIndex).Error
	fmt.Println("index: "+ followedDashboardIndex)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Failed to update components of user's followed"})
		return
	}
	// Update dashboards with the components
	updateQuery := `
		UPDATE dashboards
		SET components = ?, updated_at = CURRENT_TIMESTAMP
		WHERE "index" = '` + followedDashboardIndex + "'"
	err = models.DBManager.Exec(updateQuery, dashboard.Components).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Failed to update components of user's followed"})
		fmt.Println(http.StatusBadRequest, err)
		return
	}

	// Form the response message based on follow status
	var response string
	if isFollow {
		response = fmt.Sprintf("User %d followed Dashboard: %s", userID, dashboardIndex)
	} else {
		response = fmt.Sprintf("User %d unfollowed Dashboard: %s", userID, dashboardIndex)
	}

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"status": "success", "response": response, "follow": isFollow})
}


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

// perform insert or delete on database
func update_user_follow(userID int, index string) (bool, error) {
	var action bool // action = true: follow. Otherwise, unfollow dashboard
	table := models.DBManager.Table("user_followed")
	err := table.Raw("SELECT update_user_follow(?, ?)", userID, index).Scan(&action).Error
	if err != nil {
		return false, err
	}
	return action, nil
}
