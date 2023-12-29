// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// mathbasics.CountNumOfDigits(12345)
	// // arr := [6]int{13, 46, 24, 52, 20, 9}
	// arr := [7]int{2, 13, 4, 1, 3, 6, 28}
	// fmt.Println("WITHOUT SORT: ", arr)
	// // sorting.SelectionSort(arr)
	// sorting.BubbleSort(arr)

	// var name string = "nimai"
	// fmt.Printf("Type: %T \n", name)
	//Taking userinput from commandline
	// fmt.Println("Enter the reading for user input: ")
	// reader := bufio.NewReader(os.Stdin)

	// //comma, OK / err, OK syntax
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("You entered:", input)
	// fmt.Printf("\nYour input type: %T\n", input) //Printf needs \n at end, otherwise it wont print

	fmt.Println("Pizza Rating Application")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your rating: (1-5)")

	input, _ := reader.ReadString('\n')

	fmt.Println("Input is : ", input)
}
