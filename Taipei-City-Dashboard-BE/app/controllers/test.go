// controllers/data.go

package controllers

import (
	"TaipeiCityDashboardBE/app/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetData handles the retrieval of data
func GetData(c *gin.Context) {
	data, err := models.GetData()  // 確保有 GetData 方法在 models 中實現
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": data})
}
