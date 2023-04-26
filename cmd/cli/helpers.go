package main

import (
	"math"
	"sort"
	"time"
)

// getDateRanges returns a slice of date ranges to be used for filtering transactions.
func (app *application) getDateRanges(filterByDate bool) [][]time.Time {
	var dateRanges [][]time.Time

	if filterByDate {
		startEndDates, keys := app.generateStartEndDates()
		for _, key := range keys {
			startDate := key
			endDate := startEndDates[startDate]
			dateRanges = append(dateRanges, []time.Time{startDate, endDate})
		}
	} else {
		startDate := (*app.transactions)[len(*app.transactions)-1].Date
		endDate := (*app.transactions)[0].Date
		dateRanges = append(dateRanges, []time.Time{startDate, endDate})
	}

	return dateRanges
}

// generateStartEndDates returns a map of start and end dates for each 4 week period.
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

// calculateSavingsRate returns the savings rate for a given date range.
func (app *application) calculateSavingsRate(totalIncomes, totalExpenses float64) float64 {
	absTotalExpenses := math.Abs(totalExpenses)
	savings := totalIncomes - absTotalExpenses
	savingsRate := (savings / totalIncomes) * 100

	// Limit the maximum value of the savings rate to 100%
	if savingsRate > 100 {
		savingsRate = 100
	}

	return savingsRate
}

// filterByDateRange returns a slice of transactions that fall within the given date range.
func (app *application) filterByDateRange(startDate, endDate time.Time) Transactions {
	var filtered Transactions
	for _, transaction := range *app.transactions {
		if transaction.Date.After(startDate) && transaction.Date.Before(endDate) {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}
