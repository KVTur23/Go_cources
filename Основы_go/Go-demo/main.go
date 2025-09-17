package main

import (
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	// for i := 0; i < 10; i++ {
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Printf("%d\n", i)
	// }

	fmt.Println("__Калькулятор индекса массы тела__")
	userHeiht, userWeight := getUserInput()
	IMT := calculateIMT(userHeiht, userWeight)
	switch {
	case IMT < 16:
		fmt.Println("У Вас сильный недостаток веса")
	case IMT < 18.5:
		fmt.Println("Дефицит массы тела")
	case IMT < 25:
		fmt.Println("Нормальный вес")
	case IMT < 30:
		fmt.Println("Избыточный вес")
	default:
		fmt.Println("Степень ожирения")

	}

	// isLean := IMT < 16
	// if isLean {
	// 	fmt.Println("У Вас сильный недостаток веса")
	// } else if IMT < 18.5 {
	// 	fmt.Println("Дефицит массы тела")
	// } else if IMT < 25 {
	// 	fmt.Println("Нормальный вес")
	// } else if IMT < 30 {
	// 	fmt.Println("Избыточный вес")
	// } else {
	// 	fmt.Println("Степень ожирения")
	// }

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
