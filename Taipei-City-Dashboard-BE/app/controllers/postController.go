// controllers/postController.go
package controllers

import (
	"net/http"
	"strconv"

	"TaipeiCityDashboardBE/app/models"
	"github.com/gin-gonic/gin"
)

// LikePost 處理按讚請求
func LikePost(c *gin.Context) {
	postIDStr := c.Param("id")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	err = models.UpdateLikes(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to update likes"})
		return
	}

	c.Status(http.StatusNoContent)
}
