package main

import (
	"fmt"
	"os"
)

func main() {
	// testConsole()
	testArraySlice()
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

		// If user enter choice and press enter, one extra new line character will be there at the input buffer
		//This extra ScanLn will consume that
		fmt.Scanln()

		switch choice {
		case 1:
			fmt.Println(ini, "Creating DB")
			createFile()
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

func createFile() {
	// Take input from the user
	var fileName string
	fmt.Println("Enter your database Name: ")
	fmt.Scanln(&fileName) //assuming user will enter something
	fmt.Println("Creating database ", fileName)
	if createDatabase(fileName) {
		fmt.Println("Database created successfully")
	} else {
		panic("something went wrong") //Later add nice color sceme
	}
}

func createDatabase(fileName string) bool {
	return false
}

// testing arrays and slice concept
func testArraySlice() {
	var arr [5]int
	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Size: ", len(arr), " Capacity: ", cap(arr)) //Print size, capacity
	fmt.Println("values without initialize: ", arr)          //prints the array wit all value as 0, if string empty space
	fmt.Println("GET: ", arr[3])                             //Works, returns 0

	fmt.Println("B array: ", b, "Size: ", len(arr), " Capacity: ", cap(arr)) // Declare and initialize at the same place

	//SLICES --< An array without size

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0) // prints emty arrray, array values are nil and size is 0

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s)) // cap() -> capacity

	s[0] = "Hello"
	s[1] = "World"

	fmt.Println("After PUT:", s, "len:", len(s), "cap:", cap(s)) // cap() -> capacity

	fmt.Println("ARRAY: ", b[2:5], "Slice", s[:1]) //Both works

}
