package main

import (
	// "encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/PSUSWENG894/BudgetAPI/db"
	"github.com/PSUSWENG894/BudgetAPI/account"
	"github.com/PSUSWENG894/BudgetAPI/income"
)

var shouldInitiateDate bool

func init() {
	fmt.Printf("Running init")
	db.SetupDatabase()
}

func initiateData(database *gorm.DB){
	account.InitiateData(database)
	income.InitiateData(database)
}

func main() {
	shouldInitiateDate = true
	msg := "Hello World!"
	fmt.Printf(msg)

	router := gin.Default()
	router.GET("/", func(ctxt *gin.Context){
		ctxt.JSON(200, gin.H{"message": msg},)
	})

	apiGroup := router.Group("/api")
	database := db.GetDB()
	
	account.RegisterAccountsRoutes(apiGroup.Group("account"))
	account.Migrate(database)

	income.Migrate(database)

	if shouldInitiateDate {
		initiateData(database)
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
