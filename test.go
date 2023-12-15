// Online Go compiler to run Golang program online
// Print "Hello World!" message

package main

import "fmt"

func main() {
	arr := [6]int{2, 5, 6, 1, 3, 7}
	fmt.Println(arr)
	var a = make([]int, 6)
	for i := 0; i < len(arr); i++ {
		a = append(a, arr[i])
	}
	fmt.Println(a)
}
