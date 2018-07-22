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
	db.AutoMigrate(&Expense{})
}

func RecreateTable(db *gorm.DB) {
	db.DropTable(&Expense{})
	Migrate(db)
}

func InitiateData(db *gorm.DB) {
	var count int
	db.Model(&Expense{}).Count(&count)
	print("Expense Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialExpense())
	}
}
