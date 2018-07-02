package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/PSUSWENG894/BudgetAPI/db"
	"github.com/PSUSWENG894/BudgetAPI/model"
)


func init() {
	fmt.Printf("Running init")
	db.SetupDatabase(true)
}

func main() {

	msg := "Hello World!"
	fmt.Printf(msg)

	router := gin.Default()
	router.GET("/", func(ctxt *gin.Context){
		ctxt.JSON(200, gin.H{"message": msg},)
	})

	routes := router.Group("/api/account")
	{
		routes.POST("/", createAccount)
		routes.GET("/", fetchAllAccounts)
		routes.GET("/:id", fetchAccount)
		routes.PUT("/:id", updateAccount)
		routes.DELETE("/:id", deleteAccount)
	}



	router.Run()
}

func createAccount(ctxt *gin.Context){
	var account model.Account
	ctxt.BindJSON(&account)

	database := db.GetDB()
	database.Save(&account)

	ctxt.JSON(201, account)
}
func fetchAllAccounts(ctxt *gin.Context){
	msg := "Fetching all accounts"
	fmt.Printf(msg)

	database := db.GetDB()
	accounts := []model.Account{}
	database.Find(&accounts)
	var count int
	database.Model(&model.Account{}).Count(&count)
	
	msg = ""
	acctJson, _ := json.Marshal(accounts)
	msg += string(acctJson)

	ctxt.JSON(200, gin.H{"message": msg},)
}
func fetchAccount(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Fetching account " + id
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func updateAccount(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Updating account " + id
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func deleteAccount(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Deleting account " + id
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}