package income

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/PSUSWENG894/BudgetAPI/db"
)

func RegisterIncomeRoutes(router *gin.RouterGroup) {
	router.POST("/", createIncome)
	router.GET("/", fetchAllIncomes)
	router.GET("/:id", fetchIncome)
	router.PUT("/:id", updateIncome)
	router.DELETE("/:id", deleteIncome)
}

func createIncome(ctxt *gin.Context){
	var income Income
	ctxt.BindJSON(&income)

	database := db.GetDB()
	database.Save(&income)

	ctxt.JSON(201, income)
}
func fetchAllIncomes(ctxt *gin.Context){
	msg := "Fetching all incomes"
	fmt.Printf(msg)

	database := db.GetDB()
	incomes := []Income{}
	database.Find(&incomes)
	var count int
	database.Model(&Income{}).Count(&count)
	
	msg = ""
	incomeJson, _ := json.Marshal(incomes)
	msg += string(incomeJson)

	ctxt.JSON(200, gin.H{"message": msg},)
}
func fetchIncome(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Fetching income " + id
	fmt.Printf(msg)

	database := db.GetDB()
	income := Income{}
	database.Find(&income, id)

	msg = ""
	incomeJson, _ := json.Marshal(income)
	msg += string(incomeJson)

	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func updateIncome(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Updating income " + id

	database := db.GetDB()

	income := Income{}
	database.Find(&income, id)
	ctxt.BindJSON(&income)

	database.Save(&income)

	print("Bound income to context" + "\n")
	incomeJson, _ := json.Marshal(income)
	print(string(incomeJson) + "\n")	

	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func deleteIncome(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Deleting income " + id
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}