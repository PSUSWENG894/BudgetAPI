package account

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/PSUSWENG894/BudgetAPI/db"
	"github.com/PSUSWENG894/BudgetAPI/budget"
)

func RegisterAccountRoutes(router *gin.RouterGroup) {
	router.POST("/", createAccount)
	router.GET("/", fetchAllAccounts)
	router.GET("/:id", fetchAccount)
	router.PUT("/:id", updateAccount)
	router.DELETE("/:id", deleteAccount)
}

// Accounts
// @Summary Add Account
// @Description Create an account
// @Tags accounts account
// @Accept json
// @Produce json
// @Success 201 {object} Account
// @Router /api/account [post]
func createAccount(ctxt *gin.Context){
	var account Account
	ctxt.BindJSON(&account)

	database := db.GetDB()
	database.Save(&account)

	ctxt.JSON(201, account)
}

// Accounts
// @Summary List Accounts
// @Description Get all accounts
// @Tags accounts account
// @Produce json
// @Success 201 {array} Account
// @Router /api/account [get]
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

	database := db.GetDB()
	account := Account{}
	database.Find(&account, id)

	msg = ""
	acctJson, _ := json.Marshal(account)
	msg += string(acctJson)

	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func updateAccount(ctxt *gin.Context){
	isValid := true
	msg := ""

	id := ctxt.Params.ByName("id")
	database := db.GetDB()
	account := Account{}
	database.Find(&account, id)
	ctxt.BindJSON(&account)

	budget := budget.Budget{}
	database.Find(&budget, account.BudgetID)
	if budget.ID < 1 {
		isValid = false
	}

	if isValid {
		database.Save(&account)
		acctJson, _ := json.Marshal(account)
		msg += string(acctJson)
		ctxt.JSON(200, gin.H{"data": msg},)
	} else {
		msg = "Invalid data"
		ctxt.JSON(400, gin.H{"error": msg},)
	}
}
func deleteAccount(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Deleting account " + id + "\n"
	fmt.Printf(msg)

	database := db.GetDB()

	account := Account{}
	d := database.Delete(&account, id)
	print(d)
	msg = "Account of id " + string(id) + " deleted"

	ctxt.JSON(200, gin.H{"message": msg},)
}