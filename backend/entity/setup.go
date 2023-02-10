package entity

import (
	"time"

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
		&EmergencyLevel{},
		&HighBloodPressureLevel{},
		&DiabetesLevel{},
		&ObesityLevel{},
		&OutpatientScreening{},
	)

	db = database

	//Underscore ซ่อน err อยู่ เนื่องจากเราไม่ได้ใช้ (ไม่งั้นมันจะขึ้นแจ้งเตือนสีเหลือง) go ตัวแปรที่เขียนขึ้นมันต้องเอาไปใช้อะ
	Password, _ := bcrypt.GenerateFromPassword([]byte("123456"), 14)

	dob2, err := time.Parse("03 Feb 2536", "03 Feb 1993")
	if err != nil {
		panic(err)
	}
	db.Model(&Nurse{}).Create(&Nurse{
		FirstName:            "Chatree",
		LastName:             "Nernplab",
		Email:                "ratchanok@gmail.com",
		Password:             string(Password),
		IdentificationNumber: "1234567890123",
		BirthDay:             dob2,
		Mobile:  "0812345678",
		Address: "123/456",
		Salary:  22000,
	})
	dob4, err := time.Parse("05 Feb 2536", "05 Feb 1993")
	if err != nil {
		panic(err)
	}
	db.Model(&Nurse{}).Create(&Nurse{
		FirstName:            "Firstname",
		LastName:             "Lastname",
		Email:                "chanatip@example.com",
		Password:             string(Password),
		IdentificationNumber: "1234567890154",
		BirthDay:             dob4,
		Mobile:  "0892355678",
		Address: "143/486",
		Salary:  22000,
	})

	var ratchanok Nurse
	var chanatip Nurse
	db.Raw("SELECT * FROM nurses WHERE email = ?", "ratchanok@gmail.com").Scan(&ratchanok)
	db.Raw("SELECT * FROM nurses WHERE email = ?", "chanatip@example.com").Scan(&chanatip)

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
		Level:             "กลุ่มเสี่ยงสูงมาก",
		AssessmentForms:   "Diabetes risk score มากกว่า 8",
		HistoryTakingForm: "Text Field",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&DiabetesLevel1)

	DiabetesLevel2 := DiabetesLevel{
		Level:             "กลุ่มเสี่ยงสูง",
		AssessmentForms:   "Diabetes risk score 6-8 คะแนน",
		HistoryTakingForm: "Text Field",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&DiabetesLevel2)

	DiabetesLevel3 := DiabetesLevel{
		Level:             "กลุ่มเสี่ยงปานกลาง",
		AssessmentForms:   "Diabetes risk score 3-5 คะแนน",
		HistoryTakingForm: "Text Field",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&DiabetesLevel3)

	DiabetesLevel4 := DiabetesLevel{
		Level:             "กลุ่มปกติ",
		AssessmentForms:   "Diabetes risk score น้อยกว่า 3 คะแนน",
		HistoryTakingForm: "Text Field",
	}
	db.Model(&HighBloodPressureLevel{}).Create(&DiabetesLevel4)

	// ObesityLevel Data
	ObesityLevel1 := ObesityLevel{
		Level:             "กลุ่มผิดปกติ",
		AssessmentForms:   "BMI > 35, มีความผิดปกติมากกว่าเท่ากับ 3 ข้อจาก 5 ข้อในการซักประวัติ",
		HistoryTakingForm: "Text Field",
	}
	db.Model(&ObesityLevel{}).Create(&ObesityLevel1)

	ObesityLevel2 := ObesityLevel{
		Level:             "กลุ่มปกติ",
		AssessmentForms:   "BMI < 35, มีความผิดปกติน้อยกว่า 3 ข้อจาก 5 ข้อในการซักประวัติ",
		HistoryTakingForm: "Text Field",
	}
	db.Model(&ObesityLevel{}).Create(&ObesityLevel2)

	// OutpatienSceening Data
	OutpatientScreening1 := OutpatientScreening{
		HistorySheet:           HistorySheet1,
		EmergencyLevel:         EmergencyLevel1,
		HighBloodPressureLevel: HighBloodPressureLevel1,
		DiabetesLevel:          DiabetesLevel1,
		ObesityLevel:           ObesityLevel1,
	}
	db.Model(&OutpatientScreening{}).Create(&OutpatientScreening1)

	OutpatientScreening2 := OutpatientScreening{
		HistorySheet:           HistorySheet1,
		EmergencyLevel:         EmergencyLevel2,
		HighBloodPressureLevel: HighBloodPressureLevel2,
		DiabetesLevel:          DiabetesLevel2,
		ObesityLevel:           ObesityLevel2,
	}
	db.Model(&OutpatientScreening{}).Create(&OutpatientScreening2)
}
