package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = getInputUser("Revenue: ")
	expenses = getInputUser("Expenses: ")
	taxRate = getInputUser("Tax Rate: ")

	ebt, profit, ratio := calculatingFinancial(revenue, expenses, taxRate)

	fmt.Printf("EBT: %.1f\n", ebt)

	fmt.Printf("Profit: %.1f\n", profit)

	fmt.Printf("Ratio: %.3f\n", ratio)
}

func getInputUser(userInput string) float64 {
	var inputUser float64
	fmt.Print(userInput)
	fmt.Scan(&inputUser)
	return inputUser
}

func calculatingFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	profit := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / profit
	return earningBeforeTax, profit, ratio
}
