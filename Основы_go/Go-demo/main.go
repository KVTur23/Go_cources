package main

import (
	"errors"
	"fmt"
	"math"
)

const IMTPower = 2

func main() {
	fmt.Println("__Калькулятор индекса массы тела__")
	for {
		userHeiht, userWeight := getUserInput()
		IMT, err := calculateIMT(userHeiht, userWeight)
		if err != nil {
			// fmt.Println("Не заданы параметры для рассчета")
			// continue
			panic("Не заданы параметры для рассчета")
		}
		outputResult(IMT)
		isRepeateCalculation := checkRepeatCalculation()
		if !isRepeateCalculation {
			break
		}
	}
}

func outputResult(imt float64) {
	switch {
	case imt < 16:
		fmt.Println("У Вас сильный недостаток веса")
	case imt < 18.5:
		fmt.Println("Дефицит массы тела")
	case imt < 25:
		fmt.Println("Нормальный вес")
	case imt < 30:
		fmt.Println("Избыточный вес")
	default:
		fmt.Println("Степень ожирения")
	}
	result := fmt.Sprintf("Ваш индекс массы тела: %.0f", imt)
	fmt.Print(result)

}

func calculateIMT(userHeiht float64, userWeight float64) (float64, error) {
	if userHeiht <= 0 || userWeight <= 0 {
		return 0, errors.New("No params ERROR")
	}
	IMT := userWeight / math.Pow(userHeiht/100, IMTPower)
	return IMT, nil
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

func checkRepeatCalculation() bool {
	var userChoise string
	fmt.Print("\nВы хотите сделать ещё расчет (y/n)")
	fmt.Scan(&userChoise)
	if userChoise == "y" || userChoise == "Y" {
		return true
	}
	return false
}
