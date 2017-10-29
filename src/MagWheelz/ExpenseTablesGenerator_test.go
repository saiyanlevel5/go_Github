package main

import "testing"

func Test_GenerateTables(t *testing.T) {

	var expense *Expense = new(Expense)
	expense.MonthName = "January"
	expense.OverLimit = "20"
	expense.Saved = "10.53"
	expense.TotalExpense = "100.00"
	expense.Fuel = "45.98"
	expense.LifeStyle = "56"
	expense.Eats = "76"
	expense.Grocery = "83"
	var expenseArray []Expense
	expenseArray = append(expenseArray, *expense)
	expense.GenerateTables(&expenseArray)

}