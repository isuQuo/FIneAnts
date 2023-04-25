package main

import (
	"sort"
)

type Trend struct {
	Description string
	TotalAmount float64
}

type Trends []Trend

// Override sort.Interface methods
func (t Trends) Len() int           { return len(t) }
func (t Trends) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t Trends) Less(i, j int) bool { return t[i].TotalAmount > t[j].TotalAmount }

func (app *application) findTopTrends(topX int) Trends {
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

	sort.Sort(trends)
	if topX > 0 && len(trends) > topX {
		trends = trends[:topX]
	}

	return trends
}
