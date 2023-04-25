package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

type Trend struct {
	Description string
	TotalAmount float64
}

type Trends []Trend

// Override sort.Interface methods
func (t Trends) Len() int      { return len(t) }
func (t Trends) Swap(i, j int) { t[i], t[j] = t[j], t[i] }
func (t Trends) Less(i, j int, asc bool) bool {
	if asc {
		return t[i].TotalAmount < t[j].TotalAmount
	} else {
		return t[i].TotalAmount > t[j].TotalAmount
	}
}

func (app *application) findTopTrends(topX int, isPayment bool) Trends {
	trends := make(Trends, 0)
	descriptionAmountMap := make(map[string]float64)

	for _, transaction := range *app.transactions {
		if val, ok := descriptionAmountMap[transaction.Description]; ok {
			descriptionAmountMap[transaction.Description] = val + transaction.Amount
		} else {
			descriptionAmountMap[transaction.Description] = transaction.Amount
		}
	}

	for description, totalAmount := range descriptionAmountMap {
		trends = append(trends, Trend{
			Description: description,
			TotalAmount: totalAmount,
		})
	}

	sort.Slice(trends, func(i, j int) bool {
		if isPayment {
			return trends[i].TotalAmount > trends[j].TotalAmount
		}
		return trends[i].TotalAmount < trends[j].TotalAmount
	})

	if topX > 0 && len(trends) > topX {
		trends = trends[:topX]
	}

	return trends
}

func (app *application) generateStartEndDates() (map[time.Time]time.Time, []time.Time) {
	earliestDate := (*app.transactions)[0].Date
	latestDate := (*app.transactions)[len(*app.transactions)-1].Date

	weeks := math.Abs(latestDate.Sub(earliestDate).Hours() / 24 / 7)
	startEndDates := make(map[time.Time]time.Time)
	for i := 0; i < int(weeks); i += 4 {
		startDate := earliestDate.AddDate(0, 0, -(i+4)*7)
		endDate := earliestDate.AddDate(0, 0, -i*7)
		startEndDates[startDate] = endDate
	}

	var keys []time.Time
	for k := range startEndDates {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i].Before(keys[j])
	})

	return startEndDates, keys
}

// trends.go
func (app *application) printTopTransactionTrends(startDate, endDate time.Time, topX int, isPayment bool) {
	oldTransactions := app.transactions
	filteredTransactions := app.filterByDateRange(startDate, endDate)

	// Filter transactions based on whether they are payments or expenses
	filteredTypeTransactions := Transactions{}
	for _, transaction := range filteredTransactions {
		if (transaction.Amount >= 0) == isPayment {
			filteredTypeTransactions = append(filteredTypeTransactions, transaction)
		}
	}

	app.transactions = &filteredTypeTransactions
	topTrends := app.findTopTrends(topX, isPayment)

	if len(topTrends) > 0 {
		if isPayment {
			fmt.Printf("Top Payments Trends for %s to %s:\n", startDate.Format("02-01-2006"), endDate.Format("02-01-2006"))
		} else {
			fmt.Printf("Top Expenses Trends for %s to %s:\n", startDate.Format("02-01-2006"), endDate.Format("02-01-2006"))
		}
		var total float64
		for _, trend := range topTrends {
			fmt.Printf("  %s: $%.2f\n", trend.Description, trend.TotalAmount)
			total += trend.TotalAmount
		}
		fmt.Printf("Total: $%.2f\n", total)
		fmt.Print("\n")
	}

	app.transactions = oldTransactions
}

func (app *application) printTopTrends(topX int, filterByDate bool) {
	var keys []time.Time
	var startEndDates map[time.Time]time.Time
	if filterByDate {
		startEndDates, keys = app.generateStartEndDates()
	} else {
		keys = make([]time.Time, 1)
	}

	for _, key := range keys {
		var startDate, endDate time.Time
		if filterByDate {
			startDate = key
			endDate = startEndDates[startDate]
		} else {
			// If filterByDate is false, set startDate and endDate to the earliest and latest transaction dates
			startDate = (*app.transactions)[len(*app.transactions)-1].Date
			endDate = (*app.transactions)[0].Date
		}

		oldTransactions := app.transactions
		if filterByDate {
			filteredTransactions := app.filterByDateRange(startDate, endDate)
			app.transactions = &filteredTransactions
		}

		for _, txType := range []bool{true, false} {
			app.printTopTransactionTrends(startDate, endDate, topX, txType)
		}

		app.transactions = oldTransactions
	}
}
