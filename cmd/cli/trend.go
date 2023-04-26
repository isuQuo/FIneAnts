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

// calculateTopTrends calculates top trends across all transactions
func (app *application) calculateTopTrends(transactions Transactions, topX int, txType TransactionType) []Trend {
	trends := make(Trends, 0)
	descriptionAmountMap := make(map[string]float64)

	filteredTransactions := app.filterTransactionsByType(transactions, txType)
	for _, transaction := range filteredTransactions {
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
		if txType == Payment {
			return trends[i].TotalAmount > trends[j].TotalAmount
		}
		return trends[i].TotalAmount < trends[j].TotalAmount
	})

	if topX > 0 && len(trends) > topX {
		trends = trends[:topX]
	}

	return trends
}

// printTopTrendsByDate prints top trends for a given date range
func (app *application) printTopTrendsByDate(startDate, endDate time.Time, topTrends []Trend, txType TransactionType) {
	if len(topTrends) > 0 {
		if txType == Payment {
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
}

// printTopTrends prints top trends for a given date range
func (app *application) printTopTrends(topX int, filterByDate bool) {
	dateRanges := app.getDateRanges(filterByDate)

	for _, dates := range dateRanges {
		startDate := dates[0]
		endDate := dates[1]

		filteredTransactions := app.filterByDateRange(startDate, endDate)

		var totalIncomes, totalExpenses float64
		for _, txType := range []TransactionType{Payment, Expense} {
			filteredTypeTransactions := app.filterTransactionsByType(filteredTransactions, txType)
			topTrends := app.calculateTopTrends(filteredTypeTransactions, topX, txType)
			app.printTopTrendsByDate(startDate, endDate, topTrends, txType)

			if txType == Payment {
				for _, trend := range topTrends {
					totalIncomes += trend.TotalAmount
				}
			} else {
				for _, trend := range topTrends {
					totalExpenses += math.Abs(trend.TotalAmount)
				}
			}
		}

		savingsRate := app.calculateSavingsRate(totalIncomes, totalExpenses)
		fmt.Printf("Total: $%.2f\n", totalIncomes-totalExpenses)
		fmt.Printf("Savings Rate: %.2f%%\n", savingsRate)
		fmt.Println("--------------------------------------------------")
	}
}
