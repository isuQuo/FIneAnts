package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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
	topTrendsXPtr := flag.Int("x", 0, "number of top trends to calculate")
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

	if *topTrendsPtr != 0 {
		topTrends := app.findTopTrends(*topTrendsPtr)
		fmt.Println("Top Trends:")
		for _, trend := range topTrends {
			fmt.Printf("  %s: $%.2f\n", trend.Description, trend.TotalAmount)
		}
	}

	if *topTrendsXPtr != 0 {
		// Get earliest and latest dates
		earliestDate := transactions[0].Date
		latestDate := transactions[len(transactions)-1].Date

		// Calculate the number of weeks between the two dates and split into 4 weeks
		weeks := math.Abs(latestDate.Sub(earliestDate).Hours() / 24 / 7)
		// Create a slice of the start dates of each week every 4 weeks
		startEndDates := make(map[time.Time]time.Time)
		// This loop will create a map of start dates and end dates for each 4 week period
		for i := 0; i < int(weeks); i += 4 {
			startDate := earliestDate.AddDate(0, 0, -(i+4)*7)
			endDate := earliestDate.AddDate(0, 0, -i*7)
			startEndDates[startDate] = endDate
		}

		// Sort startEndDates by start date and loop through startDate and endDate chronologically
		var keys []time.Time
		for k := range startEndDates {
			keys = append(keys, k)
		}
		sort.Slice(keys, func(i, j int) bool {
			return keys[i].Before(keys[j])
		})

		// Filter transactions by each 4 week period and calculate the top trends
		for key := range keys {
			startDate := keys[key]
			endDate := startEndDates[startDate]

			oldTransactions := app.transactions
			filteredTransactions := app.filterByDateRange(startDate, endDate)

			app.transactions = &filteredTransactions

			topTrends := app.findTopTrends(*topTrendsXPtr)
			fmt.Printf("Top Trends for %s to %s:\n", startDate.Format("02-01-2006"), endDate.Format("02-01-2006"))
			for _, trend := range topTrends {
				fmt.Printf("  %s: $%.2f\n", trend.Description, trend.TotalAmount)
			}
			fmt.Print("\n")

			app.transactions = oldTransactions
		}
	}
}
