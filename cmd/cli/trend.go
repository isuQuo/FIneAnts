package main

import (
	"sort"
)

type Trend struct {
	Description string
	TotalAmount float64
}

type Trends []Trend

func (t Trends) Len() int           { return len(t) }
func (t Trends) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Trends) Less(i, j int) bool { return t[i].TotalAmount > t[j].TotalAmount }

func (app *application) findTopTrends(topX int) Trends {
	descriptionAmountMap := make(map[string]float64)

	for _, transaction := range *app.transactions {
		if transaction.Amount < 0 { // Only consider expenses
			descriptionAmountMap[transaction.Description] += -transaction.Amount
		}
	}

	var trends Trends
	for description, totalAmount := range descriptionAmountMap {
		trend := Trend{
			Description: description,
			TotalAmount: totalAmount,
		}
		trends = append(trends, trend)
	}

	sort.Sort(trends)

	if len(trends) > topX {
		trends = trends[:topX]
	}

	return trends
}
