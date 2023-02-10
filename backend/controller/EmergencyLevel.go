package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

/* Get คือดึงตาม id ที่ส่งไป(ส่งไปหรือส่งมาว้ะ 5555) ส่วน list คือดึงทั้งหมด*/

// POST /emergency_levels
func CreateEmergencyLevel(c *gin.Context) {
	var emergency_level entity.EmergencyLevel
	if err := c.ShouldBindJSON(&emergency_level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&emergency_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": emergency_level})
}

// GET /emergency_level/:id
func GetEmergencyLevel(c *gin.Context) {
	var emergency_level entity.EmergencyLevel

	//ใช้ Preload("Owner") หรอ?
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM food_payment_types WHERE id = ?", id).Find(&emergency_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": emergency_level})
}

// GET /emergency_levels
func ListEmergencyLevels(c *gin.Context) {
	var emergency_level []entity.EmergencyLevel

	if err := entity.DB().Raw("SELECT * FROM food_payment_types").Scan(&emergency_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": emergency_level})
}
