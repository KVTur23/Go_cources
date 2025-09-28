package main

import (
	"fmt"
)

func main() {
	tr := make([]string, 0, 2)
	tr = append(tr, "1", "2", "3")
	fmt.Println(tr)

	fmt.Println("Программа приема транзакций")
	transactions := []float64{}
	for {
		new_pay := usrInput()
		if new_pay == 0 {
			break
		}
		transactions = append(transactions, new_pay)
		if new_pay == 0 {
			break
		}
	}
	budget := summTransaction(transactions)
	fmt.Print("full budget is ", budget)
}

func usrInput() float64 {
	var pay float64
	fmt.Print("input transaction pay, or 'n' for end ")
	fmt.Scan(&pay)
	return pay
}

func summTransaction(trans []float64) float64 {
	budget := float64(0)
	for _, value := range trans {
		budget += value
	}
	return budget
}
