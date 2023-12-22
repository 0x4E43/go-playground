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

}
