package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func setupDatabase(){
	dbConnString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", 
		os.Getenv("DATABASE_URL"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"))
	db, err := sql.Open("postgres", dbConnString)
}