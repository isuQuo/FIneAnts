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
	topTrendsPtr := flag.Bool("t", false, "calculate top trends")
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

	// Test for flags and call correct function
	if *totalExpensesPtr {
		totalExpenses, totalIncome := app.calculateTotalExpensesAndIncome()
		fmt.Printf("Total Expenses: $%.2f\n", totalExpenses)
		fmt.Printf("Total Income: $%.2f\n", totalIncome)
	}

	if *topTrendsPtr {
		topTrends := app.findTopTrends(10)
		fmt.Println("Top Trends:")
		for _, trend := range topTrends {
			fmt.Printf("  %s: $%.2f\n", trend.Description, trend.TotalAmount)
		}
	}
}
