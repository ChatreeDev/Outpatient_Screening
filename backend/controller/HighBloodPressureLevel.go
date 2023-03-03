package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

/* Get คือดึงตาม id ที่ส่งไป(ส่งไปหรือส่งมาว้ะ 5555) ส่วน list คือดึงทั้งหมด*/

// POST /highbloodpressure_levels
func CreateHighBloodPressureLevel(c *gin.Context) {
	var highbloodpressure_level entity.HighBloodPressureLevel
	if err := c.ShouldBindJSON(&highbloodpressure_level); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&highbloodpressure_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": highbloodpressure_level})
}

// GET /highbloodpressure_level/:id
func GetHighBloodPressureLevel(c *gin.Context) {
	var highbloodpressure_level entity.HighBloodPressureLevel

	//ใช้ Preload("Owner") หรอ?
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM high_blood_pressure_levels WHERE id = ?", id).Scan(&highbloodpressure_level).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": highbloodpressure_level})
}

// GET /highbloodpressure_levels
func ListHighBloodPressureLevels(c *gin.Context) {
	var highbloodpressure_levels []entity.HighBloodPressureLevel

	if err := entity.DB().Raw("SELECT * FROM high_blood_pressure_levels").Scan(&highbloodpressure_levels).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": highbloodpressure_levels})
}

/* Note.
1. preload จะใช้ก็ต่อเมื่อ ตารางหลักต้องการดึงรายละเอียดต่างๆ(ข้อมูล)ของตารางรองไปใส่
ก็คือปกติมันจะดึงแค่ตัว ID ไปแต่มันจะไม่ดึงพวก object เราเขียน preload เพื่อดึงรายละเอียด
มันไปด้วย ex. ตาราง Listbookings ที่มีการ preload("member")
2. เมื่อเราใส่ preload จะใส่ scan ไม่ได้ ใส่ได้แค่ find
3. ชื่อที่ใช้ใน database มันจะใช้เป็นพหูพจน์ (เติม s)
4.
*/
