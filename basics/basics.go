package basics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func IOOperation() {
	fmt.Println("Pizza Rating Application")
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your rating: (1-5)")

	input, _ := reader.ReadString('\n')

	fmt.Println("Thank you for rating, ", input)

	//Trim string before conversion

	input = strings.TrimSpace(input)
	numRating, err := strconv.ParseFloat(input, 64) //error will come because \n is added to input

	if err != nil {
		fmt.Println("Error: ", err)
	} else {
		fmt.Println("Adding +1 ", numRating+1)
	}
}
