package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// 定義新的結構來存儲查詢結果
type ComponentLikes struct {
	ComponentID int64 `json:"component_id" gorm:"component_id"`
	Likes       int64 `json:"likes" gorm:"total_likes"`
}

// 定義 Dashboard 結構
type Dashboard struct {
	ID         int           `json:"-" gorm:"column:id;autoincrement;primaryKey"`
	Index      string        `json:"index" gorm:"column:index;type:varchar;unique;not null"`
	Name       string        `json:"name" gorm:"column:name;type:varchar;not null"`
	Components pq.Int64Array `json:"components" gorm:"column:components;type:int[]"`
	Icon       string        `json:"icon" gorm:"column:icon;type:varchar;not null"`
	UpdatedAt  time.Time     `json:"updated_at" gorm:"column:updated_at;type:timestamp with time zone;not null"`
	CreatedAt  time.Time     `json:"-" gorm:"column:created_at;type:timestamp with time zone;not null"`
}

// 定義返回的數據結構
type AllDashboards struct {
	Public   []Dashboard `json:"public"`
	Personal []Dashboard `json:"personal"`
}

type Response struct {
	Data   AllDashboards `json:"data"`
	Status string        `json:"status"`
}

// IsLike 函数检查指定用户是否已经对指定组件点过赞
func IsLike(userID int, componentID int) (bool, error) {
	var count int64

	// 执行 SQL 查询以检查记录是否存在
	query := `
        SELECT COUNT(*)
        FROM users_like
        WHERE user_id = ? AND component_id = ?
    `
	result := models.DBManager.Raw(query, userID, componentID).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}

	// 如果 count 大于 0，则表示用户已经点过赞
	return count > 0, nil
}

// IsLikeHandler 处理检查用户是否已经点过赞的请求
func IsLikeHandler(c *gin.Context) {
	// 获取用户 ID 并转换为 int64
	userID := c.GetInt("accountID")

	// 获取组件 ID
	componentIDStr := c.Param("componentid")
	componentID, er := strconv.Atoi(componentIDStr)
	fmt.Println(componentIDStr)
	fmt.Println(componentID)
	if er != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check like status 1"})
		return
	}
	// 调用 IsLike 函数检查点赞状态
	isLiked, err := IsLike(userID, componentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to check like status"})
		return
	}

	// 返回结果
	c.JSON(http.StatusOK, gin.H{"is_liked": isLiked})
}

func LikeComponentByID(c *gin.Context) {

	// Get the user ID from the context
	userID := c.GetInt("accountID")
	// Get the component ID from the form
	componentIDstr := c.Param("componentid")
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

// FetchPostsOrderByLikes 獲取排序後的 Dashboard
func FetchPostsOrderByLikes() (Response, error) {
	var componentLikes []ComponentLikes
	var response Response
	var dashboard Dashboard

	// 編寫 SQL 查詢來計算每個 component_id 的喜歡數並排序
	query := `
        SELECT component_id, COUNT(*) AS total_likes
        FROM users_like
        GROUP BY component_id
        ORDER BY total_likes DESC
    `

	// 使用 GORM 的 Raw 方法來執行查詢
	result := models.DBManager.Raw(query).Scan(&componentLikes)
	if result.Error != nil {
		return response, result.Error
	}

	// 填充 Dashboard 結構的 Components 欄位
	for _, cl := range componentLikes {
		dashboard.Components = append(dashboard.Components, cl.ComponentID)
	}

	// 更新 Dashboards 表中的 components 欄位
	updateQuery := `
		UPDATE dashboards
		SET components = ?
		WHERE index = 'likes-components'
	`
	updateResult := models.DBManager.Exec(updateQuery, dashboard.Components)
	if updateResult.Error != nil {
		return response, updateResult.Error
	}

	// 構建返回的數據結構
	response.Data = AllDashboards{
		Public: []Dashboard{
			{
				Index:      "likes-components",
				Name:       "熱度排行",
				Components: dashboard.Components,
				Icon:       "bug_report",
				UpdatedAt:  time.Now(),
			},
			{
				Index:      "map-layers",
				Name:       "圖資資訊",
				Components: dashboard.Components,
				Icon:       "public",
				UpdatedAt:  time.Now(),
			},
		},
		Personal: []Dashboard{}, // 如果有個人化的Dashboard數據，可以在這裡填充
	}
	response.Status = "success"

	return response, nil
}
