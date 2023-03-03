package entity

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Employee struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Password  string
	Salary    uint16

	// 1 Employee สามารถลงบันทึกได้หลาย HistorySheet
	HistorySheets []HistorySheet `gorm:"foreignKey:EmployeeID"`
}

type HistorySheet struct {
	gorm.Model
	Weight                 float32
	Height                 float32
	BMI                    float32
	Temperature            float32
	SystolicBloodPressure  uint16
	DiastolicBloodPressure uint16
	HeartRate              uint8
	RespiratoryRate        uint8
	OxygenSaturation       uint8
	DrugAllergySymtom      string
	PatientSymtom          string

	// 1 HistorySheet สามารถมีการประเมินได้หลายครั้ง (OutpatientScreening)
	OutpatientScreenings []OutpatientScreening `gorm:"foreignKey:HistorySheetID"`

	// EmployeeID คือ Foreign Key ที่อ้างอิงไปยัง Employee
	EmployeeID *uint
	Employee   Employee `gorm:"references:ID"`
}

// EmergencyLevel
type EmergencyLevel struct {
	gorm.Model

	Level           string
	AssessmentForms string
	procedure       string

	OutpatientScreenings []OutpatientScreening `gorm:"foreignKey:EmergencyLevelID"`
}

type HighBloodPressureLevel struct {
	gorm.Model

	Level           string
	AssessmentForms string

	OutpatientScreenings []OutpatientScreening `gorm:"foreignKey:HighBloodPressureLevelID"`
}
type DiabetesLevel struct {
	gorm.Model

	Level           string
	AssessmentForms string

	OutpatientScreenings []OutpatientScreening `gorm:"foreignKey:DiabetesLevelID"`
}
type ObesityLevel struct {
	gorm.Model

	Level           string
	AssessmentForms string

	OutpatientScreenings []OutpatientScreening `gorm:"foreignKey:ObesityLevelID"`
}

type OutpatientScreening struct {
	gorm.Model
	Note      string    `valid:"required~กรุณากรอกการซักประวัติเพิ่มเติม"`
	Date      time.Time `valid:"present~Date must be present"`
	TimeStart time.Time `valid:"future~Start Time must be future"`
	TimeEnd   time.Time `valid:"future~End Time must be future"`

	HistorySheetID *uint
	HistorySheet   HistorySheet `gorm:"references:ID"`

	EmergencyLevelID *uint
	EmergencyLevel   EmergencyLevel `gorm:"references:ID" valid:"-"`

	HighBloodPressureLevelID *uint
	HighBloodPressureLevel   HighBloodPressureLevel `gorm:"references:ID" valid:"-"`

	DiabetesLevelID *uint
	DiabetesLevel   DiabetesLevel `gorm:"references:ID" valid:"-"`

	ObesityLevelID *uint
	ObesityLevel   ObesityLevel `gorm:"references:ID" valid:"-"`
}

func init() {
	govalidator.CustomTypeTagMap.Set("past", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.Before(time.Now())
	})
	govalidator.CustomTypeTagMap.Set("future", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now())
	})
	govalidator.CustomTypeTagMap.Set("present", func(i interface{}, context interface{}) bool {
		t := i.(time.Time)
		return t.After(time.Now().Add(time.Hour*-12)) && t.Before(time.Now().Add(time.Hour*12))
	})
}
