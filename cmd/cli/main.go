package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

type application struct {
	errorLog     *log.Logger
	infoLog      *log.Logger
	transactions *Transactions
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	filenamePtr := flag.String("f", "", "filename to process")
	totalExpensesPtr := flag.Bool("e", false, "calculate total expenses")
	topTrendsPtr := flag.Int("t", 0, "calculate top trends")
	topTrendsXPtr := flag.Int("x", 0, "number of top trends to calculate")
	excludeTransactionsPtr := flag.String("ex", "", "exclude transactions with this description")
	includeTransactionsPtr := flag.String("in", "", "include transactions with this description")
	greaterAmountPtr := flag.Float64("ga", 0, "include transactions greater or equal than this amount")
	lesserAmountPtr := flag.Float64("la", 0, "include transactions less or equal than this amount")
	flag.Parse()

	if *filenamePtr == "" {
		fmt.Println("Please provide a filename using the -f flag")
		return
	}

	transactions, err := importCSV(*filenamePtr)
	if err != nil {
		errorLog.Fatalf("Unable to import CSV: %s", err)
	}

	app := &application{
		errorLog:     errorLog,
		infoLog:      infoLog,
		transactions: &transactions,
	}

	// -ex flag
	if *excludeTransactionsPtr != "" {
		transactions := app.excludeTransactions(*excludeTransactionsPtr)
		app.transactions = &transactions
	}

	// -in flag
	if *includeTransactionsPtr != "" {
		transactions := app.includeTransactions(*includeTransactionsPtr)
		app.transactions = &transactions
	}

	// -ga flag
	if *greaterAmountPtr != 0 {
		transactions := app.greaterAmount(*greaterAmountPtr)
		app.transactions = &transactions
	}

	// -la flag
	if *lesserAmountPtr != 0 {
		transactions := app.lesserAmount(*lesserAmountPtr)
		app.transactions = &transactions
	}

	// After filtering transactions, check if there are any left
	if len(*app.transactions) == 0 {
		fmt.Println("No transactions found")
		return
	}

	// -e flag
	if *totalExpensesPtr {
		totalExpenses, totalIncome := app.calculateTotalExpensesAndIncome()
		fmt.Printf("Total Expenses: $%.2f\n", totalExpenses)
		fmt.Printf("Total Income: $%.2f\n", totalIncome)
	}

	// -t flag
	if *topTrendsPtr != 0 {
		app.printTopTrends(*topTrendsPtr, false)
	}

	// -x flag
	if *topTrendsXPtr != 0 {
		app.printTopTrends(*topTrendsXPtr, true)
	}
}
