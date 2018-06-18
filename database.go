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

	createTables()
}

funct createTables(){
	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS account (name varchar(25),balance NUMERIC(17,2))"); err != nil {
        fmt.Printf("Error creating database table: account", err)
    }
}