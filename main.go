package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)


func init() {
	fmt.Printf("Running init")
	setupDatabase()
}

func testdb(ctxt *gin.Context) {
	msg := ""

	err := db.Ping()
	if err != nil {
		msg = "Got error connecting to db: " + err.Error()
	} else {
		msg = "Pinged DB successfully: " + os.Getenv("DB_HOST")
		
	}
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
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
	msg := "Creating account"
	fmt.Printf(msg)
	ctxt.JSON(200, gin.H{"message": msg},)
}
func fetchAllAccounts(ctxt *gin.Context){
	msg := "Fetching all accounts"
	fmt.Printf(msg)

	qry := `SELECT * FROM account`
	results, err := db.Query(qry)
	if err != nil {
		msg = "Error fetching accounts from database"
		fmt.Printf(msg)
	}

	type Account struct {
		Name string
		Balance float64
	}

	accounts := make([]Account,0)
	msg = ""
	for results.Next() {
		var name string
		var balance float64
		results.Scan(&name, &balance)
		acct := Account{name, balance}
		accounts = append(accounts, acct)
		acctJson, _ := json.Marshal(acct)
		msg += string(acctJson)
	}

	fmt.Printf("Total accounts: %d", len(accounts))

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