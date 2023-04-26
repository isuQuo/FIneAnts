package main

import (
	"fmt"
)

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

// -gd flag
func (app *application) handleGreaterDateFlag(date string) {
	if date != "" {
		transactions := app.filterDate(app.parseDate(date), true)
		app.transactions = &transactions
	}
}

// -ld flag
func (app *application) handleLesserDateFlag(date string) {
	if date != "" {
		transactions := app.filterDate(app.parseDate(date), false)
		app.transactions = &transactions
	}
}

// -md flag
func (app *application) handleMiddleDateFlag(dates []string) {
	if len(dates) == 2 {
		startDate := app.parseDate(dates[0])
		endDate := app.parseDate(dates[1])
		transactions := app.filterByDateRange(startDate, endDate)
		app.transactions = &transactions
	}
}
