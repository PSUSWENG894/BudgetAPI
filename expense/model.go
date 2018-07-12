package expense

import (
	"github.com/jinzhu/gorm"
)

type Expense struct {
	gorm.Model
	Name string `json:"name"`
	Amount float64 `json:"amount"`
}

func GetInitialExpense() *Expense {
	expns := Expense{Name: "Work", Amount: 231.36}
	return &expns
}

func Migrate(db *gorm.DB) {
	print("Migrating expense")
	db.AutoMigrate(&Expense{})
}

func InitiateData(db *gorm.DB) {
	var count int
	db.Model(&Expense{}).Count(&count)
	print("Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialExpense())
	}
}