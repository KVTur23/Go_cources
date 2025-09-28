package main

import "fmt"

type stringMap = map[string]string

func main() {
	bookmarks := stringMap{}
	fmt.Println("Bookmarks CHOOSER")
Menu:
	for {
		choice := menu()
		switch choice {
		case 1:
			for key, value := range bookmarks {
				fmt.Println(key, ":", value)
			}
		case 2:
			addBookmarks(bookmarks)
		case 3:
			removeBookmarks(bookmarks)
		case 4:
			fmt.Println("\n Exiting from programm..")
			break Menu
		}
	}
}

func menu() int {
	var choice int
	fmt.Print("\nChoose your choice:" +
		" \n- 1. look the bookmarks\n- 2. add bookmarks" +
		"\n- 3. remove bookmarks\n -4. exit\n")
	fmt.Scan(&choice)
	return choice
}

func addBookmarks(bookmarks stringMap) {
	var newKey, newValue string
	fmt.Println("Please, input a new bookmark name")
	fmt.Scan(&newKey)
	fmt.Println("Please, input a new bookmark URL")
	fmt.Scan(&newValue)
	bookmarks[newKey] = newValue
}

func removeBookmarks(bookmarks stringMap) {
	var delKey string
	fmt.Println("Please, input the  bookmark name for REMOVE")
	fmt.Scan(&delKey)
	delete(bookmarks, delKey)
}
