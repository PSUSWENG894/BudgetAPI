package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/PSUSWENG894/BudgetAPI/db"
)

func RegisterAccountRoutes(router *gin.RouterGroup) {
	router.POST("/", createAccount)
	router.GET("/", fetchAllAccounts)
	router.GET("/:id", fetchAccount)
	router.PUT("/:id", updateAccount)
	router.DELETE("/:id", deleteAccount)
}

func createAccount(ctxt *gin.Context){
	var account Account
	ctxt.BindJSON(&account)

	database := db.GetDB()
	database.Save(&account)

	ctxt.JSON(201, account)
}
func fetchAllAccounts(ctxt *gin.Context){
	msg := "Fetching all accounts"
	fmt.Printf(msg)

	database := db.GetDB()
	accounts := []Account{}
	database.Find(&accounts)
	var count int
	database.Model(&Account{}).Count(&count)
	
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