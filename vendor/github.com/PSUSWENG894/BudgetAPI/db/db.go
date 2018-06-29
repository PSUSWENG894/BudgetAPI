package db

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
	"os"
)

var db *gorm.DB

type Account struct {
	gorm.Model
	Name string `json:"name"`
	Balance float64 `json:"balance"`
}

func SetupDatabase(initiateWithData bool){
	msg := ""
	dbConnString := fmt.Sprintf("host=%s dbname=%s user=%s password=%s sslmode=disable", 
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), 
		os.Getenv("DB_PASSWORD"))
	// database, err := sql.Open("postgres", dbConnString)
	database, err := gorm.Open("postgres", dbConnString)
	db = database
	if err != nil {
		msg += ". Got error connecting to db: " + err.Error()
		fmt.Printf(msg)
		return
	}

	migrate()
	if initiateWithData {
		initiateData()
	}
}

func migrate() {
	db.AutoMigrate(&Account{})
}

func initiateData() {
	var count int
	db.Model(&Account{}).Count(&count)
	print("Count: ", count, "\n")
	if count == 0 {
		db.Create(&Account{Name: "test", Balance: 0.00})
	}
}

func GetDB() *gorm.DB {
	return db
}
