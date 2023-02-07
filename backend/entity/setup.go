package entity

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema
	database.AutoMigrate(
		&Nurse{},
		&HistorySheet{},
		&HighBloodPressureLevel{},
		&DiabetesLevel{},
		&ObesityLevel{},
		&OutpatientScreening{},
	)

	db = database

	//Underscore ซ่อน err อยู่ เนื่องจากเราไม่ได้ใช้ (ไม่งั้นมันจะขึ้นแจ้งเตือนสีเหลือง) go ตัวแปรที่เขียนขึ้นมันต้องเอาไปใช้อะ
	Password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//ใส่ Data ยังไม่ครบ
	db.Model(&Nurse{}).Create(&Nurse{
		FirstName: "Chatree",
		LastName:  "Nernplab",
		//Email:     "chatree@gmail.com",
		Password: string(Password),
	})
	db.Model(&Nurse{}).Create(&Nurse{
		FirstName: "Firstname",
		LastName:  "Lastname",
		Email:     "name@example.com",
		Password:  string(Password),
	})

	var chatree Nurse
	var name Nurse
	db.Raw("SELECT * FROM users WHERE email = ?", "chatree@gmail.com").Scan(&chatree)
	db.Raw("SELECT * FROM users WHERE email = ?", "name@example.com").Scan(&name)

	// --- FoodSet Data
	SnackBox := DiabetesLevel{
		Level:           "ปกติ",
		AssessmentForms: "ไม่ต้องประเมินผล",
	}
	db.Model(&DiabetesLevel{}).Create(&SnackBox)

	CoffeeBreak := DiabetesLevel{
		Level:           "ปกติ",
		AssessmentForms: "ไม่ต้องประเมินผล",
	}
	db.Model(&DiabetesLevel{}).Create(&CoffeeBreak)

	PremiumBakery := DiabetesLevel{
		Level:           "ปกติ",
		AssessmentForms: "ไม่ต้องประเมินผล",
	}
	db.Model(&DiabetesLevel{}).Create(&PremiumBakery)

	// FoodPaymentType Data
	PromptPay := HighBloodPressureLevel{
		Level:           "ปกติ",
		AssessmentForms: "ไม่ต้องประเมินผล",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&PromptPay)

	BankTransfer := HighBloodPressureLevel{
		Level:           "ปกติ",
		AssessmentForms: "ไม่ต้องประเมินผล",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&BankTransfer)

	Cash := HighBloodPressureLevel{
		Level:           "ปกติ",
		AssessmentForms: "ไม่ต้องประเมินผล",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&Cash)

	//เหลือ Data ของ Booking ใส่ไม่เป็นอะ
	HistorySheet1 := HistorySheet{
		Weight: 61.6,
		Height: 171.0,
	}
	db.Model(&HistorySheet{}).Create(&HistorySheet1)

	ObesityLevel1 := ObesityLevel{
		Level:           "ปกติ",
		AssessmentForms: "ไม่ต้องประเมินผล",
	}
	db.Model(&ObesityLevel{}).Create(&ObesityLevel1)

	// FoodOrdered Data
	OutpatientScreening1 := OutpatientScreening{
		HistorySheet:           HistorySheet1,
		HighBloodPressureLevel: PromptPay,
		DiabetesLevel:          SnackBox,
		ObesityLevel:           ObesityLevel1,
	}
	db.Model(&OutpatientScreening{}).Create(&OutpatientScreening1)

}
