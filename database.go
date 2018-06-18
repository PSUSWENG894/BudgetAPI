package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var db *sql.DB

func setupDatabase(){
	msg := ""
	dbConnString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", 
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"))
	database, err := sql.Open("postgres", dbConnString)
	db = database
	if err != nil {
		msg += ". Got error connecting to db: " + err.Error()
		fmt.Printf(msg)
		return
	}
	err = db.Ping()
	if err != nil {
		msg += ". Got error connecting to db: " + err.Error()
		fmt.Printf(msg)
		return
	} else {
		fmt.Printf("Pinged DB successfully: " + os.Getenv("DATABASE_URL"))
	}

	createTables()
}

func createTables(){
	tableName := "account"
	tableDefinition := tableName + ` (name varchar(25) UNIQUE NOT NULL,
		balance NUMERIC(17,2))
		`
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS " + tableDefinition); err != nil {
        fmt.Printf("Error creating database table %s: %s", tableName, err)
        return
    }else {
    	fmt.Printf("Created database table %s", tableName)
    }

    insertStatement := `
    INSERT INTO account(name, balance) values ('test', 0.00)`
    _, err := db.Exec(insertStatement)
    if err != nil {
    	fmt.Printf("Error add account")
    }
}