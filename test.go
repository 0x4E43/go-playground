// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func main() {
	arr := [6]int{2, 5, 6, 1, 3, 0}
	fmt.Println(arr)
	// var a = make([]int, 6)
	for i := 0; i < len(arr)-1; i++ {
		small := arr[i]
		for j := i; j < len(arr); j++ {
			if arr[j] <= small {
				small = arr[j]
				arr[j] = arr[i]
				arr[i] = small

			}
		}
	}
	fmt.Println(arr)

}
