package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /DiabetesLevels
func CreateDiabetesLevel(c *gin.Context) {
	var diabetesLevel entity.DiabetesLevel
	if err := c.ShouldBindJSON(&diabetesLevel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&diabetesLevel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": diabetesLevel})
}

// GET /DiabetesLevel/:id
func GetDiabetesLevel(c *gin.Context) {
	var diabetesLevel entity.DiabetesLevel

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM food_sets WHERE id = ?", id).Find(&diabetesLevel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diabetesLevel})
}

// GET /DiabetesLevels
func ListDiabetesLevels(c *gin.Context) {
	var diabetesLevels []entity.DiabetesLevel //[] ส่งเป็นแบบลิสต์

	if err := entity.DB().Raw("SELECT * FROM food_sets").Scan(&diabetesLevels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": diabetesLevels})
}

// DELETE /DiabetesLevels/:id
func DeleteDiabetesLevel(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM food_sets WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food set not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /DiabetesLevels
func UpdateDiabetesLevel(c *gin.Context) {
	var DiabetesLevel entity.DiabetesLevel
	if err := c.ShouldBindJSON(&DiabetesLevel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", DiabetesLevel.ID).First(&DiabetesLevel); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food set not found"})
		return
	}

	if err := entity.DB().Save(&DiabetesLevel).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": DiabetesLevel})
}

/* Note.
1. DiabetesLevel เป็นตารางรอง ไม่จำเป็นต้องมี preload เพราะไม่ต้องไปดึงของใครมาใส่ของตัวเอง
2.
*/
