package sorting

import "fmt"

// TODO: will use this later
// type Array struct {
// 	arr []int
// }

func SelectionSort(arr [7]int) {
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
	fmt.Println("SELECTION SORT: ", arr)
}

func BubbleSort(arr [7]int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			small := arr[j]
			if arr[i] > arr[j] {
				arr[i] = arr[j]
				arr[j] = small
			}
		}
	}
	fmt.Println("BUBBLE SORT: ", arr)
}
