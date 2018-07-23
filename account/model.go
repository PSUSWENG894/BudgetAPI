package account

import (
	"github.com/jinzhu/gorm"
	"github.com/PSUSWENG894/BudgetAPI/budget"
)

type Account struct {
	gorm.Model
	Name string `json:"name"`
	Balance float64 `json:"balance"`
	IncludeInBudget bool `json"includeInBudget"`
	Budget budget.Budget `json"budget"`
	BudgetID int `json"budgetID"`
}

func GetInitialAccount() *Account {
	acct := Account{Name: "TestAccount", Balance: 7.00, IncludeInBudget: true}
	return &acct
}

func Migrate(db *gorm.DB) {
	print("Creating table")
	print(&Account{})
	db.AutoMigrate(&Account{})
}

func RecreateTable(db *gorm.DB) {
	print("Dropping table accounts")
	db.DropTable("accounts")
	print("Dropped table accounts")
	Migrate(db)
}

func InitiateData(db *gorm.DB) {
	var count int
	db.Model(&Account{}).Count(&count)
	print("Acount Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialAccount())
	}
}
