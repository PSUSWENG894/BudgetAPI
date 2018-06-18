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

	// createTables()
}

// funct createTables(){
// 	if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
//         fmt.Printf("Error creating database table: %q", err)
//     }
// }