package main

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

/*
	based on the supplied data for the expense, do string to int conversion and calcualte data for the above three
once you have the code running, refactor it around the way each month value is set
 */
type ExpenseTablesGenerator interface {

	GenerateTables(expenses *[]Expense)

}


type Expense struct {
	MonthName string
	Grocery string
	LifeStyle string
	Eats string
	Fuel string
	Saved string
	OverLimit string
	TotalExpense string
}


func (Expense) GenerateTables(expenses *[]Expense) {

	/*var MonthlyExpenseCollection [][] string = [][]string {
		{
			expenses.month, expenses.Grocery, expenses.LifeStyle, expenses.Eats, expenses.Fuel, expenses.Totalexpenses, expenses.Saved, expenses.OverLimit,
		},
	}*/

	var MonthlyExpenseCollection [][]string

	for _, expense := range *expenses {
		var expenseInString []string
		expenseInString = append(expenseInString, expense.MonthName)
		expenseInString = append(expenseInString, expense.Grocery)
		expenseInString = append(expenseInString, expense.LifeStyle)
		expenseInString = append(expenseInString, expense.Eats)
		expenseInString = append(expenseInString, expense.Fuel)
		expenseInString = append(expenseInString, expense.TotalExpense)
		expenseInString = append(expenseInString, expense.Saved)
		expenseInString = append(expenseInString, expense.OverLimit)
		LOGGER.Debug("Entries in Expense String Array: %v", expenseInString)
		MonthlyExpenseCollection = append(MonthlyExpenseCollection, expenseInString)
		LOGGER.Debug("MonthExpenseCollection So Far: %v", MonthlyExpenseCollection)
		expenseInString = nil
	}

	var expenseTable *tablewriter.Table = tablewriter.NewWriter(os.Stdout)
	expenseTable.SetHeader([]string{"Month", "Grocery", "LifeStyle", "Eats", "Fuel", "Total Expense", "Saved", "OverLimit"})
	expenseTable.SetFooter([]string{"", "", "", "", "", "", "Total", ""})
	expenseTable.SetBorder(false)
	expenseTable.SetHeaderColor(
		tablewriter.Color(tablewriter.Bold, tablewriter.BgGreenColor),
		tablewriter.Color(tablewriter.FgHiRedColor, tablewriter.Bold, tablewriter.BgBlackColor),
		tablewriter.Color(tablewriter.BgRedColor, tablewriter.FgWhiteColor),
		tablewriter.Color(tablewriter.BgCyanColor, tablewriter.FgWhiteColor),
		tablewriter.Color(tablewriter.BgGreenColor, tablewriter.FgWhiteColor),
		tablewriter.Color(tablewriter.BgBlueColor, tablewriter.FgWhiteColor),
		tablewriter.Color(tablewriter.BgHiBlueColor, tablewriter.FgWhiteColor),
		tablewriter.Color(tablewriter.BgHiMagentaColor, tablewriter.FgWhiteColor),
	)

	expenseTable.SetColumnColor(
		tablewriter.Color(tablewriter.Bold, tablewriter.FgHiBlackColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.FgHiRedColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.FgBlackColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.FgBlackColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.FgBlackColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.FgBlackColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.FgBlackColor),
		tablewriter.Color(tablewriter.Bold, tablewriter.FgBlackColor),
	)

	expenseTable.SetFooterColor(
		tablewriter.Color(tablewriter.Bold),
		tablewriter.Color(tablewriter.Bold),
		tablewriter.Color(tablewriter.Bold),
		tablewriter.Color(tablewriter.Bold),
		tablewriter.Color(tablewriter.Bold),
		tablewriter.Color(tablewriter.Bold),
		tablewriter.Color(tablewriter.Bold),
		tablewriter.Color(tablewriter.FgHiRedColor))

	expenseTable.AppendBulk(MonthlyExpenseCollection)
	expenseTable.Render()

}


