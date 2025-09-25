package main

import (
	"fmt"
)

func main() {
	fmt.Println("Программа приема транзакций")
	transactions := []float64{}
	for {
		new_pay := usrInput()
		if new_pay == 0 {
			break
		}
		transactions = append(transactions, new_pay)
		repeat := checkRepeatTransaction()
		if !repeat {
			break
		}
	}
	budget := summTransaction(transactions)
	fmt.Print("full budget is ", budget)
}

func usrInput() float64 {
	var pay float64
	fmt.Print("input transaction pay ")
	fmt.Scan(&pay)
	//if pay <= 0 {
	//	return 0, errors.New("No params ERROR")
	//}
	return pay
}
func checkRepeatTransaction() bool {
	var userChoise string
	fmt.Print("\nNeed some transactions? (y/n)")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}

func summTransaction(trans []float64) float64 {
	budget := float64(0)
	for i := 0; i < len(trans); i++ {
		budget += trans[i]
	}
	return budget
}
