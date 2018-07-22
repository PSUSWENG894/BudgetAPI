package budget

import (
	"github.com/jinzhu/gorm"
)

type Budget struct {
	gorm.Model
	Name string `json:"name"`
}

func GetInitialBudget() *Budget {
	budget := Budget{Name: "Current Budget"}
	return &budget
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Budget{})
}

func InitiateData(db *gorm.DB) {
	var count int
	db.Model(&Budget{}).Count(&count)
	print("Budget Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialBudget())
	}
}