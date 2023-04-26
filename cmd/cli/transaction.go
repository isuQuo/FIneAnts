package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
	"time"
)

type TransactionType string

const (
	Income  TransactionType = "Income"
	Expense TransactionType = "Expense"
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
			transactionType = Income
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

// calculateTotalExpensesAndIncome calculates total expenses and income
func (app *application) calculateTotalExpensesAndIncome() (float64, float64) {
	var totalExpenses, totalIncome float64
	for _, transaction := range *app.transactions {
		if transaction.Type == Income {
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

func (app *application) filterTransactionsByType(transactions Transactions, txType TransactionType) Transactions {
	filtered := Transactions{}
	for _, transaction := range transactions {
		if transaction.Type == txType {
			filtered = append(filtered, transaction)
		}
	}
	return filtered
}
