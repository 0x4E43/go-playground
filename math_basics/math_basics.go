package mathbasics

import (
	"fmt"
	"strconv"
)

func CountNumOfDigits(num int) {
	digits := 0
	for i := num; i != 0; i = i / 10 {
		digits++
	}
	fmt.Println("Number of digits in : ", num, " : ", digits)

	//method 2 -> Convert into string and calculate String length

	len := len(strconv.Itoa(num))
	fmt.Println("Number of digits in(using string method) : ", num, " : ", len)
	countNumberOfEvenlyDividingDigits(num)
}

// PROBLEM LINK: https://www.codingninjas.com/studio/problems/count-digits_8416387
func countNumberOfEvenlyDividingDigits(num int) {
	if num != 0 {
		count := 0
		for i := num; i != 0; i = i / 10 {
			div := i % 10
			if div != 0 && num%div == 0 {
				count++
				fmt.Println("DIGIT: ", div)
			}
		}
		fmt.Println("Num of evenly deviding digits: ", count)
	}
}
