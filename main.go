package main

import (
	// "encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/PSUSWENG894/BudgetAPI/db"
	"github.com/PSUSWENG894/BudgetAPI/account"
)


func init() {
	fmt.Printf("Running init")
	db.SetupDatabase()
}

func main() {

	msg := "Hello Forrest!"
	fmt.Printf(msg)

	router := gin.Default()
	router.GET("/", func(ctxt *gin.Context){
		ctxt.JSON(200, gin.H{"message": msg},)
	})

	apiGroup := router.Group("/api")
	account.RegisterAccountsRoutes(apiGroup.Group("account"))

	database := db.GetDB()
	account.Migrate(database)
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
