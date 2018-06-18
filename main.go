package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB

func init() {
	msg := ""
	dbConnString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", 
		os.Getenv("DATABASE_URL"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", dbConnString)
	if err != nil {
		msg += ". Got error connecting to db: " + err.Error()
		fmt.Printf(msg)
	}
	err = db.Ping()
	if err != nil {
		msg += ". Got error connecting to db: " + err.Error()
		fmt.Printf(msg)
	} else {
		fmt.Printf("Pinged DB successfully: " + os.Getenv("DATABASE_URL"))
	}
}

func testdb(ctxt *gin.Context) {
	msg := ""

	err := db.Ping()
	if err != nil {
		msg = "Got error connecting to db: " + err.Error()
	} else {
		msg = "Pinged DB successfully: " + os.Getenv("DATABASE_URL")
		
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