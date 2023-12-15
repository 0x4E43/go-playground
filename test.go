// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func main() {
	arr := [6]int{2, 5, 6, 1, 3, 7}
	fmt.Println(arr)
	var a = make([]int, 6)
	for i := 0; i < len(arr); i++ {
		// a = append(a, arr[i]) // while using make() its initializing to slice with value 0
		a[i] = arr[i] 
	}
	fmt.Println(a)

}
