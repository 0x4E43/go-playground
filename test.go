// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func main() {
	arr := [6]int{13, 46, 24, 52, 20, 9}
	fmt.Println(arr)
	selectionSort(arr)
}

func selectionSort(arr [6]int) {
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
