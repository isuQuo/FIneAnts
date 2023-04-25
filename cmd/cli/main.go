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
	topTrendsXPtr := flag.Int("tx", 0, "number of top trends to calculate")

	excludeTransactionsPtr := flag.String("ex", "", "exclude transactions with this description")
	includeTransactionsPtr := flag.String("in", "", "include transactions with this description")

	greaterAmountPtr := flag.Float64("ga", 0, "include transactions greater or equal than this amount")
	lesserAmountPtr := flag.Float64("la", 0, "include transactions less or equal than this amount")

	flag.Parse()

	if *filenamePtr == "" {
		fmt.Println("Please provide a filename using the -f flag")
		return
	}

	if *includeTransactionsPtr != "" && *excludeTransactionsPtr != "" {
		fmt.Println("Please provide only one of -ex or -in flags")
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

	app.handleExcludeTransactionsFlag(*excludeTransactionsPtr)
	app.handleIncludeTransactionsFlag(*includeTransactionsPtr)
	app.handleGreaterAmountFlag(*greaterAmountPtr)
	app.handleLesserAmountFlag(*lesserAmountPtr)

	// After filtering transactions, check if there are any left
	if len(*app.transactions) == 0 {
		fmt.Println("No transactions found")
		return
	}

	if *totalExpensesPtr {
		app.handleTotalExpensesFlag()
	}
	app.handleTopTrendsFlag(*topTrendsPtr, false)
	app.handleTopTrendsXFlag(*topTrendsXPtr, true)
}
