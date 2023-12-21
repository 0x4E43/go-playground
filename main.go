// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import (
	"fmt"

	"github.com/0x4E43/go-playground/sorting"
)

func main() {

	arr := [6]int{13, 46, 24, 52, 20, 9}
	fmt.Println("WITHOUT SORT: ", arr)
	sorting.SelectionSort(arr)
}
