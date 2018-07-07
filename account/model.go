package account

import (
	"github.com/jinzhu/gorm"
)

type Account struct {
	gorm.Model
	Name string `json:"name"`
	Balance float64 `json:"balance"`
}

func GetInitialAccount() *Account {
	acct := Account{Name: "test", Balance: 0.00}
	return &acct
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Account{})
}

func InitiateData(db *gorm.DB) {
	var count int
	db.Model(&Account{}).Count(&count)
	print("Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialAccount())
	}
}