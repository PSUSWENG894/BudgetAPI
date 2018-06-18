package main

import (
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

	router.GET("/testdb", testdb)

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