package expense

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/PSUSWENG894/BudgetAPI/db"
)

func RegisterExpenseRoutes(router *gin.RouterGroup) {
	router.POST("/", createExpense)
	router.GET("/", fetchAllExpenses)
	router.GET("/:id", fetchExpense)
	router.PUT("/:id", updateExpense)
	router.DELETE("/:id", deleteExpense)
}

func createExpense(ctxt *gin.Context){
	var expense Expense
	ctxt.BindJSON(&expense)

	database := db.GetDB()
	database.Save(&expense)

	ctxt.JSON(201, expense)
}
func fetchAllExpenses(ctxt *gin.Context){
	msg := "Fetching all expenses"
	fmt.Printf(msg)

	database := db.GetDB()
	expenses := []Expense{}
	database.Find(&expenses)
	var count int
	database.Model(&Expense{}).Count(&count)
	
	msg = ""
	expnsJson, _ := json.Marshal(expenses)
	msg += string(expnsJson)

	ctxt.JSON(200, gin.H{"message": msg},)
}
func fetchExpense(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Fetching expense " + id
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func updateExpense(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Updating expense " + id
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func deleteExpense(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Deleting expense " + id
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}