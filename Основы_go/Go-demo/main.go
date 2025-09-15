package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	userHeiht, userWeight := getUserInput()
	IMT := calculateIMT(userHeiht, userWeight)
	outputResult(IMT)
}

func outputResult(imt float64) {
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)

}

func calculateIMT(userHeiht float64, userWeight float64) (IMT float64) {
	IMT = userWeight / math.Pow(userHeiht/100, IMTPower)
	return
}

func getUserInput() (float64, float64) {
	var userHeiht float64
	var userWeight float64
	fmt.Print("Введите Ваш рост в сантиметрах: ")
	fmt.Scan(&userHeiht)
	fmt.Print("Введите Ваш вес: ")
	fmt.Scan(&userWeight)
	return userHeiht, userWeight
}
