package controller

import (
	"github.com/ChatreeDev/sa-65-example/entity"

	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	//"github.com/asaskevich/govalidator"
	

)

// GET /Nurses
// List all Nurses
func ListNurses(c *gin.Context) {
	var nurses []entity.Nurse
	if err := entity.DB().Raw("SELECT * FROM Nurses").Scan(&nurses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nurses})
}

// GET /Nurse/:id
// Get Nurse by id
func GetNurse(c *gin.Context) {
	var nurse entity.Nurse
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM Nurses WHERE id = ?", id).Scan(&nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nurse})
}

// POST /Nurses
func CreateNurse(c *gin.Context) {
	var nurse entity.Nurse
	if err := c.ShouldBindJSON(&nurse); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// เข้ารหัสลับรหัสผ่านที่ผู้ใช้กรอกก่อนบันทึกลงฐานข้อมูล
	bytes, err := bcrypt.GenerateFromPassword([]byte(nurse.Password), 14)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error hashing password"})
		return
	}
	nurse.Password = string(bytes)

	// แทรกการ validate ไว้ช่วงนี้ของ controller
	// if _, err := govalidator.ValidateStruct(nurse); err != nil {
	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// return
	// }

	if err := entity.DB().Create(&nurse).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": nurse})
}









// PATCH /Nurses
// func UpdateNurse(c *gin.Context) {
// 	var nurse entity.Nurse
// 	if err := c.ShouldBindJSON(&nurse); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	if tx := entity.DB().Where("id = ?", nurse.ID).First(&nurse); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Nurse not found"})
// 		return
// 	}

// 	if err := entity.DB().Save(&nurse).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"data": nurse})
// }

// // DELETE /Nurses/:id
// func DeleteNurse(c *gin.Context) {
// 	id := c.Param("id")
// 	if tx := entity.DB().Exec("DELETE FROM Nurses WHERE id = ?", id); tx.RowsAffected == 0 {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Nurse not found"})
// 		return
// 	}
// 	/*
// 		if err := entity.DB().Where("id = ?", id).Delete(&entity.Nurse{}).Error; err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}*/

// 	c.JSON(http.StatusOK, gin.H{"data": id})
// }

//Create
