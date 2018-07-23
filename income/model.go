package income

import (
	"github.com/jinzhu/gorm"
	"github.com/PSUSWENG894/BudgetAPI/account"
	"time"
)

//Incomes are used while creating a budget.
//They show an estimated amount of pay to be paid at intervals defined by recurrence
type Income struct {
	gorm.Model
	Name string `json:"name"`
	Amount float64 `json:"amount"`
	Account account.Account `json"account"`
	AccountID uint `json"accountID"`
}

//Income Payments are instances of income to be received.
//These are what show up when you are looking at your budget during a time period.
type IncomePayment struct {
	gorm.Model
	Amount float64 `json:"amount"`
	PayDate time.Time `json:"payDate"`
	Income Income `json"income"`
	IncomeID uint `json"incomeID"`
}

func GetInitialIncome() *Income {
	incm := Income{Name: "Work", Amount: 1000.00}
	return &incm
}

func GetInitialIncomePayment(incomeId uint) *IncomePayment {
	dateFormat := "2006-Jan-02"
	dateString := "2018-Jul-22"
	dt := time.Now()
	dt, _ = time.Parse(dateFormat, dateString)
	incpay := IncomePayment{Amount: 1000.00, PayDate: dt, IncomeID: incomeId}
	return &incpay
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&Income{})
	db.AutoMigrate(&IncomePayment{})
}

func InitiateData(db *gorm.DB) {
	var count int
	db.Model(&Income{}).Count(&count)
	print("Income Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialIncome())
	}
	income := Income{}
	db.First(&income)

	db.Model(&IncomePayment{}).Count(&count)
	print("IncomePayment Count: ", count, "\n")
	if count == 0 {
		db.Create(GetInitialIncomePayment(income.ID))
	}
}