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

	expenses := GetExpensesFromDB()
	msg = GetExpenseListAsJson(&expenses)

	ctxt.JSON(200, gin.H{"message": msg},)
}
func GetExpensesFromDB() []Expense {
	database := db.GetDB()
	expenses := []Expense{}
	database.Find(&expenses)
	return expenses
}

func fetchExpense(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Fetching expense " + id
	fmt.Printf(msg)

	expense := getExpenseFromDB(id)
	msg = GetExpenseAsJson(&expense)

	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func getExpenseFromDB(id string) Expense {
	database := db.GetDB()
	expense := Expense{}
	database.Find(&expense, id)
	return expense
}

func updateExpense(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Updating expense " + id

	database := db.GetDB()

	expense := Expense{}
	database.Find(&expense, id)
	ctxt.BindJSON(&expense)
	database.Save(&expense)

	msg = GetExpenseAsJson(&expense)

	ctxt.JSON(200, gin.H{"message": msg},)
}
func deleteExpense(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Deleting expense " + id + "\n"
	fmt.Printf(msg)

	database := db.GetDB()

	expense := Expense{}
	database.Delete(&expense, id)
	msg = "Expense of id " + string(id) + " deleted"

	ctxt.JSON(200, gin.H{"message": msg},)
}

func GetExpenseAsJson(expense *Expense) string{
	msg := ""
	expnsJson, _ := json.Marshal(expense)
	msg += string(expnsJson)
	return msg
}
func GetExpenseListAsJson(expenseList *[]Expense) string{
	msg := ""
	expnsJson, _ := json.Marshal(expenseList)
	msg += string(expnsJson)
	return msg
}