package main

import (
	"fmt"
	"os"
)

func main() {
	testConsole()
}
func checkError(err error) {
	if err != nil {
		panic(err)
		//exit system
	}

}
func testConsole() {
	ini := "joker # "

	fmt.Println("Welcome to the Joker DB 🃏 \n\n ")
	showOptions()
	for { //infinite loop
		var choice int
		fmt.Scan(&choice)
		fmt.Println(ini, "Choice: ", choice)
		if choice == 4 {
			fmt.Println("Exiting System...")
			os.Exit(1)
		}

		switch choice {
		case 1:
			fmt.Println(ini, "Creating DB")
			break
		case 2:
			fmt.Println(ini, "Reading DB")
			break
		case 3:
			fmt.Println(ini, "Deleting DB")
			break
		}
		fmt.Println()
		fmt.Print("\n", ini, "Waiting for user input: ")
	}
}

func showOptions() {
	fmt.Println("1. Create DB")
	fmt.Println("2. Read DB")
	fmt.Println("3. Delete DB")
	fmt.Println("\nPress Esc to exit")
	fmt.Println()
}
