package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/PSUSWENG894/BudgetAPI/db"
	"github.com/PSUSWENG894/BudgetAPI/budget"
	"github.com/PSUSWENG894/BudgetAPI/account"
	"github.com/PSUSWENG894/BudgetAPI/income"
	"github.com/PSUSWENG894/BudgetAPI/expense"
	// "os"
)

var shouldInitiateData bool

func init() {
	fmt.Printf("Running init\n")
	db.SetupDatabase()
}

func initiateData(database *gorm.DB){
	budget.InitiateData(database)
	account.InitiateData(database)
	income.InitiateData(database)
	expense.InitiateData(database)
}

func CreateTestData(database *gorm.DB) {
	

	var count int
	database.Model(&budget.Budget{}).Count(&count)
	if count == 0 {
		budget1 := budget.GetInitialBudget()
		database.Create(&budget1)
		account1 := account.GetInitialAccount()
		account1.Budget = *budget1
		database.Save(&account1)	
	}

	income1 := income.Income{}
	database.Model(&income.Income{}).Count(&count)
	if count == 0 {
		income1 = *income.GetInitialIncome()
		database.Create(&income1)
	} else{
		database.First(&income1)
	}

	database.Model(&income.IncomePayment{}).Count(&count)
	if count == 0 {
		incpay := income.GetInitialIncomePayment(income1.ID)
		database.Create(&incpay)
	}	

	database.Model(&expense.Expense{}).Count(&count)
	if count == 0 {
		expense1 := expense.GetInitialExpense()
		database.Create(&expense1)
	}
}

func main() {
	migrateFlagPtr := flag.Bool("migrate", false, "Migrate database flag")
	flag.Parse()
	if *migrateFlagPtr == true {
		MigrateDatabase()
	} else {
		print("Don't migrate\n")
	}

	shouldInitiateData = true
	msg := "Hello World!"
	fmt.Printf(msg)

	router := gin.Default()
	router.GET("/", func(ctxt *gin.Context){
		ctxt.JSON(200, gin.H{"message": msg},)
	})

	apiGroup := router.Group("/api")
	database := db.GetDB()
	
	budget.RegisterBudgetRoutes(apiGroup.Group("budget"))
	budget.Migrate(database)

	account.RegisterAccountRoutes(apiGroup.Group("account"))
	account.Migrate(database)

	income.RegisterIncomeRoutes(apiGroup.Group("income"))
	income.Migrate(database)

	expense.RegisterExpenseRoutes(apiGroup.Group("expense"))
	expense.Migrate(database)

	if shouldInitiateData {
		CreateTestData(database)
	}

	// routes := router.Group("/api/account")
	// {
	// 	routes.POST("/", createAccount)
	// 	routes.GET("/", fetchAllAccounts)
	// 	routes.GET("/:id", fetchAccount)
	// 	routes.PUT("/:id", updateAccount)
	// 	routes.DELETE("/:id", deleteAccount)
	// }

	router.Run()
}
