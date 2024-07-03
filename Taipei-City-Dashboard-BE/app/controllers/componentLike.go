package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"TaipeiCityDashboardBE/app/models"

	"github.com/gin-gonic/gin"
)

// 定義新的結構來存儲查詢結果
type ComponentLikes struct {
	ComponentID int64 `json:"component_id" gorm:"component_id"`
	Likes       int64 `json:"likes" gorm:"total_likes"`
}

func LikeComponentByID(c *gin.Context) {
	// Get the user ID from the context
	userID := c.GetInt("accountID")
	// Get the component ID from the form
	componentIDstr := c.PostForm("componentid")
	fmt.Println(userID, " Like commponent ", componentIDstr, " <3<3<3")

	// turn string to int
	componentid, err := strconv.Atoi(componentIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid component ID: " + componentIDstr})
		return
	}

	// Database modify
	likeOrNot, er := UpdateLikes(userID, componentid)
	if er != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Component ID: " + componentIDstr + " not found"})
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

// GetPostsOrderByLikes 處理按照喜歡數排序並返回結果的請求
func GetPostsOrderByLikes(c *gin.Context) {
	posts, err := FetchPostsOrderByLikes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to fetch posts"})
		return
	}

	c.JSON(http.StatusOK, posts)
}

// FetchPostsOrderByLikes 統計 users_like 中每個 component_id 的總數，並按降序排列
func FetchPostsOrderByLikes() ([]ComponentLikes, error) {
	var likes []ComponentLikes

	// 編寫 SQL 查詢來計算每個 component_id 的喜歡數並排序
	query := `
        SELECT component_id, COUNT(*) AS likes
        FROM users_like
        GROUP BY component_id
        ORDER BY likes DESC
    `

	// 使用 GORM 的 Raw 方法來執行查詢
	result := models.DBManager.Raw(query).Scan(&likes)
	return likes, result.Error
}
