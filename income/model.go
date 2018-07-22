package income

import (
	"github.com/jinzhu/gorm"
)

type Income struct {
	gorm.Model
	Name string `json:"name"`
	Amount float64 `json:"amount"`
}

func GetInitialIncome() *Income {
	incm := Income{Name: "Work", Amount: 1000.00}
	return &incm
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Income{})
}

func InitiateData(db *gorm.DB) {
	var count int
	db.Model(&Income{}).Count(&count)
	print("Income Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialIncome())
	}
}