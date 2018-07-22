package main

import (
	"github.com/PSUSWENG894/BudgetAPI/db"
	"github.com/PSUSWENG894/BudgetAPI/account"
)

func MigrateDatabase() {
	print("Migrating databases\n")
	db.SetupDatabase()
	database := db.GetDB()
	account.RecreateTable(database)
	print("Done migrating databases\n")
	CreateTestData(database)
}

