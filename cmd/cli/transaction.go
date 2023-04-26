package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"
)

type TransactionType float64

const (
	// iota is a special Go constant that is used to generate a set of related but distinct constants.
	Payment TransactionType = iota
	Expense
)

type Transaction struct {
	Date        time.Time
	Amount      float64
	Description string
	Type        TransactionType
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
		var transactionType TransactionType
		if amount < 0 {
			transactionType = Expense
		} else {
			transactionType = Payment
		}

		transaction := Transaction{
			Date:        date,
			Amount:      amount,
			Description: description,
			Type:        transactionType,
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
		if transaction.Type == Payment {
			totalIncome += transaction.Amount
		} else {
			totalExpenses += -transaction.Amount
		}
	}

	return totalExpenses, totalIncome
}

// Filter transactions by excluding or including a search term
func (app *application) filterTransactions(description string, include bool) Transactions {
	var filtered Transactions
	searchTerms := strings.Split(description, "|")

	for _, transaction := range *app.transactions {
		matched := false
		lowerDesc := strings.ToLower(transaction.Description)
		for _, term := range searchTerms {
			lowerTerm := strings.ToLower(term)
			if strings.Contains(lowerDesc, lowerTerm) {
				matched = true
				break
			}
		}
		if include == matched {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}

// Filter transactions by amount
func (app *application) filterAmount(amount float64, greater bool) Transactions {
	var filtered Transactions
	for _, transaction := range *app.transactions {
		if (greater && transaction.Amount >= amount) || (!greater && transaction.Amount <= amount) {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}

// Filter transactions by date
func (app *application) filterDate(date time.Time, after bool) Transactions {
	var filtered Transactions
	for _, transaction := range *app.transactions {
		if (after && transaction.Date.After(date)) || (!after && transaction.Date.Before(date)) {
			filtered = append(filtered, transaction)
		}
	}

	return filtered
}
