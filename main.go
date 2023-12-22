// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import (
	"fmt"

	mathbasics "github.com/0x4E43/go-playground/math_basics"
	"github.com/0x4E43/go-playground/sorting"
)

func main() {

	mathbasics.CountNumOfDigits(12345)
	// arr := [6]int{13, 46, 24, 52, 20, 9}
	arr := [7]int{2, 13, 4, 1, 3, 6, 28}
	fmt.Println("WITHOUT SORT: ", arr)
	// sorting.SelectionSort(arr)
	sorting.BubbleSort(arr)
}
