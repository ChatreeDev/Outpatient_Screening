package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// GET /HistorySheets
// List all HistorySheets
func ListHistorySheets(c *gin.Context) {
	var historySheets []entity.HistorySheet
	if err := entity.DB().Raw("SELECT * FROM history_sheets").Preload("Member").Find(&historySheets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": historySheets})
}


// GET /HistorySheet/:id
// Get HistorySheet by id
func GetHistorySheet(c *gin.Context) {
	var historySheet entity.HistorySheet
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM history_sheets WHERE id = ?", id).Preload("Member").Find(&historySheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": historySheet})
}

// POST /HistorySheets
func CreateHistorySheet(c *gin.Context) {
	var HistorySheet entity.HistorySheet
	if err := c.ShouldBindJSON(&HistorySheet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&HistorySheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": HistorySheet})
}

