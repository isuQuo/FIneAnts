package main

import (
	"flag"
	"log"
	"os"
	"strings"
	"time"
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

	greaterDatePtr := flag.String("gd", "", "include transactions greater or equal than this date")
	lesserDatePtr := flag.String("ld", "", "include transactions less or equal than this date")
	middleDatePtr := flag.String("md", "", "include transactions between this date.\nSeparate dates by comma")

	flag.Parse()

	if *filenamePtr == "" {
		errorLog.Fatalln("Please provide a filename using the -f flag")
	}

	if *includeTransactionsPtr != "" && *excludeTransactionsPtr != "" {
		errorLog.Fatalln("Please provide only one of -ex or -in flags")
	}

	var dates []string
	if middleDatePtr != nil && *middleDatePtr != "" {
		if !strings.Contains(*middleDatePtr, ",") {
			errorLog.Fatalln("Please provide two dates separated by comma")
		}
		dates = strings.Split(*middleDatePtr, ",")
		if len(dates) != 2 {
			errorLog.Fatalln("Please provide two dates separated by comma")
		}
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
	app.handleGreaterDateFlag(*greaterDatePtr)
	app.handleLesserDateFlag(*lesserDatePtr)
	app.handleMiddleDateFlag(dates)

	// After filtering transactions, check if there are any left
	if len(*app.transactions) == 0 {
		app.errorLog.Fatalln("No transactions found")
	}

	if *totalExpensesPtr {
		app.handleTotalExpensesFlag()
	}
	app.handleTopTrendsFlag(*topTrendsPtr, false)
	app.handleTopTrendsXFlag(*topTrendsXPtr, true)
}

// Parse Date
func (app *application) parseDate(date string) time.Time {
	parsedDate, err := time.Parse("02-01-2006", date)
	if err != nil {
		app.errorLog.Fatalf("Unable to parse date: %s", err)
	}

	return parsedDate
}
