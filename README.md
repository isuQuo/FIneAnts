# FineAnts
A finance and budgeting application written in Go to help with financial and budgeting planning.

## usage
```
cd cmd/cli
go run . 
  -e	calculate total expenses
  -ex string
    	exclude transactions with this description
  -f string
    	filename to process
  -ga float
    	include transactions greater or equal than this amount
  -gd string
    	include transactions greater or equal than this date
  -in string
    	include transactions with this description
  -la float
    	include transactions less or equal than this amount
  -ld string
    	include transactions less or equal than this date
  -md string
    	include transactions between this date.
    	Separate dates by comma
  -t int
    	calculate top trends
  -tx int
    	number of top trends to calculate
```

## example usage
<strong>Calculate top 10 trends, exclude search terms, filter results greater than 24-04-2022</strong>
```
go run . -f ~/Downloads/BANK.csv -t 10 -ex "UBER|AMAZON" -gd 24-04-2022
```