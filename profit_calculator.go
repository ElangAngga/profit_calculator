package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	var earningBeforeTax = revenue - expenses

	var profit = earningBeforeTax * (1 - taxRate/100)

	var ratio = earningBeforeTax / profit

	fmt.Print("EBT: ", earningBeforeTax)
	fmt.Println()

	fmt.Print("Profit: ", profit)
	fmt.Println()

	fmt.Println("Ratio: ", ratio)
	fmt.Println()

}
