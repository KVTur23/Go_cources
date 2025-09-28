package main

import (
	"fmt"
)

type account struct {
	login    string
	password string
	url      string
}

func main() {
	login := promptData("Input login: ")
	password := promptData("Input password: ")
	url := promptData("Input url: ")

	myAccount := account{login, password, url}

	outputPassword(&myAccount)
}

func promptData(prompt string) string {
	fmt.Print(prompt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputPassword(acc *account) {
	fmt.Println(acc.login, acc.password, acc.url)

}

// Reverse function
//func main() {
//	a := [4]int{1, 2, 3, 4}
//	reverse(&a)
//	fmt.Println(a)
//
//}
//
//func reverse(arr *[4]int) {
//	for index, value := range *arr {
//		(*arr)[len(*arr)-1-index] = value
//	}
//}
