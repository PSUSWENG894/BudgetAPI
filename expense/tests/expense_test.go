package expense

import (
	// "fmt"
	// "github.com/PSUSWENG894/BudgetAPI/db"
	"github.com/PSUSWENG894/BudgetAPI/expense"
	"strings"
	"testing"
)

// func TestExpenses(t *testing.T){
// 	expns1 := expense.Expense{Name: "Rent", Amount: 1200.00}
// 	expns2 := expense.Expense{Name: "Food", Amount: 234.36}
// 	db.SetupDatabase()
// 	database := db.GetDB()
// 	database.Create(&expns1)
// 	database.Create(&expns2)
	
// 	expenseList := expense.GetExpensesFromDB()
// 	fmt.Println("Got %i", len(expenseList))
// 	if len(expenseList) != 2 {
// 		t.Error("expcted 2")
// 	}
// }

func TestSerializingExpenses(t *testing.T){
	expns1 := expense.Expense{Name: "Rent", Amount: 1200.00}
	expns2 := expense.Expense{Name: "Food", Amount: 234.36}
	expenseList := make([]expense.Expense, 2)
	expenseList = append(expenseList, expns1)
	expenseList = append(expenseList, expns2)
	
	// expSerialized := expense.GetExpenseAsJson(&expns1)

	expListSerialized := expense.GetExpenseListAsJson(&expenseList)
	
	// fmt.Printf("Json %s", expListSerialized)
	if !strings.Contains(expListSerialized, "Rent") {
		t.Error("expcted to find Rent")
	}
}