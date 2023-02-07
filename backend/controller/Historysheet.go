package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// GET /HistorySheets
// List all HistorySheets
func ListHistorySheets(c *gin.Context) {
	var HistorySheets []entity.HistorySheet
	if err := entity.DB().Raw("SELECT * FROM HistorySheets").Preload("Member").Find(&HistorySheets).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": HistorySheets})
}

// GET /HistorySheet/:id
// Get HistorySheet by id
func GetHistorySheet(c *gin.Context) {
	var HistorySheet entity.HistorySheet
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM HistorySheets WHERE id = ?", id).Preload("Member").Find(&HistorySheet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": HistorySheet})
}

// // POST /HistorySheets
// func CreateHistorySheet(c *gin.Context) {
// 	var HistorySheet entity.HistorySheet
// 	if err := c.ShouldBindJSON(&HistorySheet); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if err := entity.DB().Create(&HistorySheet).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"data": HistorySheet})
// }

// // PATCH /HistorySheets
// func UpdateHistorySheet(c *gin.Context) {
// 	var HistorySheet entity.HistorySheet
// 	if err := c.ShouldBindJSON(&HistorySheet); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entity.DB().Where("id = ?", HistorySheet.ID).First(&HistorySheet); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "HistorySheet not found"})
// 		return
// 	}

// 	if err := entity.DB().Save(&HistorySheet).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": HistorySheet})
// }

// // DELETE /HistorySheets/:id
// func DeleteHistorySheet(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM HistorySheets WHERE id = ?", id); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
// 		return
// 	}
// 	/*
// 		if err := entity.DB().Where("id = ?", id).Delete(&entity.User{}).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}*/

// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }
