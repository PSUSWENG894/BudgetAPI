package budget

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/PSUSWENG894/BudgetAPI/db"
)

func RegisterBudgetRoutes(router *gin.RouterGroup) {
	router.POST("/", createBudget)
	router.GET("/", fetchAllBudgets)
	router.GET("/:id", fetchBudget)
	router.PUT("/:id", updateBudget)
	router.DELETE("/:id", deleteBudget)
}

func createBudget(ctxt *gin.Context){
	var budget Budget
	ctxt.BindJSON(&budget)

	database := db.GetDB()
	database.Save(&budget)

	ctxt.JSON(201, budget)
}
func fetchAllBudgets(ctxt *gin.Context){
	msg := "Fetching all budgets"
	fmt.Printf(msg)

	database := db.GetDB()
	budgets := []Budget{}
	database.Find(&budgets)
	var count int
	database.Model(&Budget{}).Count(&count)
	
	msg = ""
	acctJson, _ := json.Marshal(budgets)
	msg += string(acctJson)

	ctxt.JSON(200, gin.H{"message": msg},)
}
func fetchBudget(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Fetching budget " + id
	fmt.Printf(msg)

	database := db.GetDB()
	budget := Budget{}
	database.Find(&budget, id)

	msg = ""
	acctJson, _ := json.Marshal(budget)
	msg += string(acctJson)

	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func updateBudget(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Updating budget " + id

	database := db.GetDB()

	budget := Budget{}
	database.Find(&budget, id)
	ctxt.BindJSON(&budget)
	database.Save(&budget)

	acctJson, _ := json.Marshal(budget)
	msg += string(acctJson)

	ctxt.JSON(200, gin.H{"message": msg},)
}
func deleteBudget(ctxt *gin.Context){
	id := ctxt.Params.ByName("id")
	msg := "Deleting budget " + id + "\n"
	fmt.Printf(msg)

	database := db.GetDB()

	budget := Budget{}
	d := database.Delete(&budget, id)
	print(d)
	msg = "Budget of id " + string(id) + " deleted"

	ctxt.JSON(200, gin.H{"message": msg},)
}