package basics

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func TimeHandling() {
	fmt.Println("Hello timeMatchine!")

	currTime := time.Now()

	fmt.Println("Current time : ", currTime) //Current time :  2024-01-03 23:39:55.539479841 +0530 IST m=+0.000091693 --? badly formatted

	//formating time
	fmt.Println("Formatted Time: ", currTime.Format("01-02-2006")) //The day go developed

	fmt.Println("Formatted Time With Day: ", currTime.Format("01-02-2006 Monday")) //The day go developed

	fmt.Println("Formatted Time With Day and Hours: ", currTime.Format("01-02-2006 Monday 15:04:05")) //The day go developed

	dummyDate := time.Date(1999, time.June, 14, 12, 5, 0, 0, time.UTC)

	fmt.Println("Dummy Date: ", dummyDate)
	fmt.Println("Formatted Dummy Date: ", dummyDate.Format("01-02-2006 Monday"))
}

func MemoryManagement() {
	//Go mamnages memory by its own
	//If anything required
	// new Vs make()
	//new -> allocate memory but do not initialize
	// make()-> does both
	// both will give a memory address
	// new -> zeroed storage
	//make -> non zeroed storage

}
