package mathbasics

import "fmt"

func CountNumOfDigits(num int) {
	digits := 0
	for i := num; i != 0; i = i / 10 {
		digits++
	}
	fmt.Println("Number of digits in : ", num, " : ", digits)
}
