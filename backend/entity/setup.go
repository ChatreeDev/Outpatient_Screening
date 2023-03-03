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
		&Employee{},
		&HistorySheet{},
		&EmergencyLevel{},
		&HighBloodPressureLevel{},
		&DiabetesLevel{},
		&ObesityLevel{},
		&OutpatientScreening{},
	)

	db = database

	//Underscore ซ่อน err อยู่ เนื่องจากเราไม่ได้ใช้ (ไม่งั้นมันจะขึ้นแจ้งเตือนสีเหลือง) go ตัวแปรที่เขียนขึ้นมันต้องเอาไปใช้อะ
	Password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	//Day of birth
	//Birth1 := "03 Feb 2005"
	// dob2, err := time.Parse("09-29-1995", "09-29-1995")
	// if err != nil {
	// 	panic(err)
	// }
	db.Model(&Employee{}).Create(&Employee{
		FirstName: "Ratchanok",
		LastName:  "Inthanon",
		Email:     "ratchanok@gmail.com",
		Password:  string(Password),
		Salary:    22000,
	})

	// dob4, err := time.Parse("09-29-1995", "09-29-1995")
	// if err != nil {
	// 	panic(err)
	// }
	db.Model(&Employee{}).Create(&Employee{
		FirstName: "Firstname",
		LastName:  "Lastname",
		Email:     "chanatip@example.com",
		Password:  string(Password),
		Salary:    22000,
	})

	var ratchanok Employee
	var chanatip Employee
	db.Raw("SELECT * FROM Employees WHERE email = ?", "ratchanok@gmail.com").Scan(&ratchanok)
	db.Raw("SELECT * FROM Employees WHERE email = ?", "chanatip@example.com").Scan(&chanatip)

	// HistorySheet Data
	HistorySheet1 := HistorySheet{
		Weight:                 61.6,
		Height:                 171.0,
		BMI:                    21.2,
		Temperature:            36.5,
		SystolicBloodPressure:  120,
		DiastolicBloodPressure: 80,
		HeartRate:              60,
		RespiratoryRate:        16,
		OxygenSaturation:       98,
		DrugAllergySymtom:      "ไม่มี",
		PatientSymtom:          "ไม่มี",
	}

	//EmergencyLevel Data
	EmergencyLevel1 := EmergencyLevel{
		Level:           "ภาวะฉุกเฉิน",
		AssessmentForms: "ESI-2 วิกฤต (Emergency) : SBP มีค่าต่ำกว่าที่ระบุ + ผิวซีดและผิวหนังเย็น, ผิวลายเป็นจ้ำๆ",
		procedure:       "แจ้งพยาบาล-แพทย์",
	}
	db.Model(&EmergencyLevel{}).Create(&EmergencyLevel1)

	EmergencyLevel2 := EmergencyLevel{
		Level:           "ภาวะเร่งด่วน",
		AssessmentForms: "ESI-3 เร่งด่วน (Urgency): HR, RR, Temp มากกว่าที่ระบุ, O2Sat < 92 %",
		procedure:       "แจ้งพยาบาล-แพทย์",
	}
	db.Model(&EmergencyLevel{}).Create(&EmergencyLevel2)

	EmergencyLevel3 := EmergencyLevel{
		Level:           "ภาวะไม่รุนแรง",
		AssessmentForms: "ESI-4  ไม่รุนแรง (Semi-Urgency): Temp มากกว่าที่ระบุ",
		procedure:       "จัดลำดับในการเข้าตรวจ",
	}
	db.Model(&EmergencyLevel{}).Create(&EmergencyLevel3)

	// HighBloodPressureLevel Data
	HighBloodPressureLevel1 := HighBloodPressureLevel{
		Level:           "กลุ่มป่วย",
		AssessmentForms: "SBP = 140 - 179 และ DBP = 90 - 109 มม.ปรอท",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&HighBloodPressureLevel1)

	HighBloodPressureLevel2 := HighBloodPressureLevel{
		Level:           "กลุ่มเสี่ยง",
		AssessmentForms: "SBP = 120 - 139 และ DBP = 80 - 89 มม.ปรอท",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&HighBloodPressureLevel2)

	HighBloodPressureLevel3 := HighBloodPressureLevel{
		Level:           "กลุ่มปกติ",
		AssessmentForms: "SBP < 120 และ DBP < 80 มม.ปรอท",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&HighBloodPressureLevel3)

	// DiabetesLevel Data
	DiabetesLevel1 := DiabetesLevel{
		Level:           "กลุ่มเสี่ยงสูงมาก",
		AssessmentForms: "Diabetes risk score มากกว่า 8",
	}
	db.Model(&DiabetesLevel{}).Create(&DiabetesLevel1)

	DiabetesLevel2 := DiabetesLevel{
		Level:           "กลุ่มเสี่ยงสูง",
		AssessmentForms: "Diabetes risk score 6-8 คะแนน",
	}
	db.Model(&DiabetesLevel{}).Create(&DiabetesLevel2)

	DiabetesLevel3 := DiabetesLevel{
		Level:           "กลุ่มเสี่ยงปานกลาง",
		AssessmentForms: "Diabetes risk score 3-5 คะแนน",
	}
	db.Model(&DiabetesLevel{}).Create(&DiabetesLevel3)

	DiabetesLevel4 := DiabetesLevel{
		Level:           "กลุ่มปกติ",
		AssessmentForms: "Diabetes risk score น้อยกว่า 3 คะแนน",
	}
	db.Model(&DiabetesLevel{}).Create(&DiabetesLevel4)

	// ObesityLevel Data
	ObesityLevel1 := ObesityLevel{
		Level:           "กลุ่มผิดปกติ",
		AssessmentForms: "BMI > 35, มีความผิดปกติมากกว่าเท่ากับ 3 ข้อจาก 5 ข้อในการซักประวัติ",
	}
	db.Model(&ObesityLevel{}).Create(&ObesityLevel1)

	ObesityLevel2 := ObesityLevel{
		Level:           "กลุ่มปกติ",
		AssessmentForms: "BMI < 35, มีความผิดปกติน้อยกว่า 3 ข้อจาก 5 ข้อในการซักประวัติ",
	}
	db.Model(&ObesityLevel{}).Create(&ObesityLevel2)

	// OutpatienSceening Data
	OutpatientScreening1 := OutpatientScreening{
		HistorySheet:           HistorySheet1,
		EmergencyLevel:         EmergencyLevel1,
		HighBloodPressureLevel: HighBloodPressureLevel1,
		DiabetesLevel:          DiabetesLevel1,
		ObesityLevel:           ObesityLevel1,
		Note:                   "คนไข้มีญาติที่เป็นโรคเบาหวาน",
		// TimeStart: 			time.Now(),
		// TimeEnd: 			time.Now(),
	}
	db.Model(&OutpatientScreening{}).Create(&OutpatientScreening1)
}
