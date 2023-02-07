package controller

import (
	"net/http"

	"github.com/ChatreeDev/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

// POST /foodordereds
func CreateOutpatientScreening(c *gin.Context) {

	var OutpatientScreening entity.OutpatientScreening
	var HistorySheet entity.HistorySheet
	// var foodset entity.FoodSet
	var HighBloodPressure_Level entity.HighBloodPressureLevel

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร foodordered
	if err := c.ShouldBindJSON(&OutpatientScreening); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา booking ด้วย id
	if tx := entity.DB().Where("id = ?", OutpatientScreening.HistorySheetID).First(&HistorySheet); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "HistorySheet not found"})
		return
	}

	// 10: ค้นหา foodset ด้วย id
	for _, orderFoodSet := range OutpatientScreening.ObesityLevel.OutpatientScreenings {
		if tx := entity.DB().Where("id = ?", orderFoodSet.DiabetesLevelID).First(&orderFoodSet); tx.RowsAffected == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "food set not found"})
			return
		}
	}

	// 11: ค้นหา foodpayment_type ด้วย id
	if tx := entity.DB().Where("id = ?", OutpatientScreening.HighBloodPressureLevelID).First(&HighBloodPressure_Level); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "food payment type not found"})
		return
	}
	// 12: สร้าง FoodOrdered
	wv := entity.OutpatientScreening{
		HistorySheet:           HistorySheet,                     // โยงความสัมพันธ์กับ Entity Booking
		HighBloodPressureLevel: HighBloodPressure_Level,          // โยงความสัมพันธ์กับ Entity FoodPaymentType
		ObesityLevel:           OutpatientScreening.ObesityLevel, // โยงความสัมพันธ์กับ Entity FoodSet (แต่ไม่โดยตรง เพราะเป็นคสพแบบหลาย)
	}

	// 13: บันทึก
	if err := entity.DB().Create(&wv).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": wv})
}

// GET /foodordereds
func ListOutpatientScreenings(c *gin.Context) {
	var OutpatientScreening []entity.OutpatientScreening

	/*เงื่อนไขสำหรับการค้นหา โดยดึงข้อมูลจากตารางรองที่เกี่ยวข้องมา #ระวัง ชื่อ field ต้องตรงกัน
	ซึ่งดูฟิลด์ได้จากเราสร้างไว้ให้ entity หลัก ในไฟล์ schema */

	if err := entity.DB().Raw("SELECT * FROM food_ordereds").
		Preload("HistorySheet").Preload("HighBloodPressureLevel").
		Preload("FoodOrderedFoodSets").Preload("FoodOrderedFoodSets.FoodSet"). //preload แบบ join table
		Find(&OutpatientScreening).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": OutpatientScreening})

}

// GET /foodordered/:id
func GetOutpatientScreening(c *gin.Context) {
	var OutpatientScreening entity.OutpatientScreening //GET จะ​ get มาแค่ก้อนเดียวเลยไม่ใช้ array (เก็ทไอดีของตัวที่เคยบันทึก) [ex. เก็ทเอาไปคิดราคา(ของระบบอื่น)]

	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM food_ordereds WHERE id = ?", id).
		Preload("HistorySheet").Preload("HighBloodPressureLevel").
		Preload("FoodOrderedFoodSets").Preload("FoodOrderedFoodSets.FoodSet").
		Find(&OutpatientScreening).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": OutpatientScreening})

}

