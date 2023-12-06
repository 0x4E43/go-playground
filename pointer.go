package main

import "fmt"

func zeroPointer(val int) {
	val = 1
}

func pointer(val *int) {
	val2 := 34
	// *val = 0
	val = &val2 //It only changes the local copy of the pointer
	fmt.Println(val2, *val)
}
func main1() {
	val := 25
	fmt.Println("Hello Pointer", val)
	zeroPointer(val)
	fmt.Println("Hello Pointer after calling by value", val)
	pointer(&val)
	fmt.Println("Hello Pointer after calling by ref", val, " Address: ", &val)

}
