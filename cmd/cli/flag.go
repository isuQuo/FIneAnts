package main

import "fmt"

// -ex flag
func (app *application) handleExcludeTransactionsFlag(description string) {
	if description != "" {
		transactions := app.filterTransactions(description, false)
		app.transactions = &transactions
	}
}

// -in flag
func (app *application) handleIncludeTransactionsFlag(description string) {
	if description != "" {
		transactions := app.filterTransactions(description, true)
		app.transactions = &transactions
	}
}

// -ga flag
func (app *application) handleGreaterAmountFlag(amount float64) {
	if amount != 0 {
		transactions := app.filterAmount(amount, true)
		app.transactions = &transactions
	}
}

// -la flag
func (app *application) handleLesserAmountFlag(amount float64) {
	if amount != 0 {
		transactions := app.filterAmount(amount, false)
		app.transactions = &transactions
	}
}

// -e flag
func (app *application) handleTotalExpensesFlag() {
	totalExpenses, totalIncome := app.calculateTotalExpensesAndIncome()
	fmt.Printf("Total Expenses: $%.2f\n", totalExpenses)
	fmt.Printf("Total Income: $%.2f\n", totalIncome)
}

// -t flag
func (app *application) handleTopTrendsFlag(numberOfTrends int, filterByDate bool) {
	if numberOfTrends != 0 {
		app.printTopTrends(numberOfTrends, false)
	}
}

// -tx flag
func (app *application) handleTopTrendsXFlag(numberOfTrends int, filterByDate bool) {
	if numberOfTrends != 0 {
		app.printTopTrends(numberOfTrends, true)
	}
}
