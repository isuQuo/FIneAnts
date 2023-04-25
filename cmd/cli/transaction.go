package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"
)

type Transaction struct {
	Date        time.Time
	Amount      float64
	Description string
}

type Transactions []Transaction

func importCSV(filename string) (Transactions, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var transactions Transactions
	for _, record := range records {
		date, _ := time.Parse("02/01/2006", record[0])
		amount, _ := strconv.ParseFloat(record[1], 64)
		description := record[2]

		transaction := Transaction{
			Date:        date,
			Amount:      amount,
			Description: description,
		}

		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (app *application) filterByDateRange(startDate, endDate time.Time) Transactions {
	var filtered Transactions
	for _, transaction := range *app.transactions {
		if transaction.Date.After(startDate) && transaction.Date.Before(endDate) {
			filtered = append(filtered, transaction)
		}
	}
	return filtered
}

func (app *application) calculateTotalExpensesAndIncome() (float64, float64) {
	var totalExpenses, totalIncome float64
	for _, transaction := range *app.transactions {
		if transaction.Amount < 0 {
			totalExpenses += -transaction.Amount
		} else {
			totalIncome += transaction.Amount
		}
	}
	return totalExpenses, totalIncome
}
