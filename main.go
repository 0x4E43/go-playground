package main

import (
	"fmt"

	"github.com/0x4E43/go-playground/basics"
	mathbasics "github.com/0x4E43/go-playground/math_basics"
)

func main() {
	// Building aplication
	// go build
	// Extended: GOOS="linux" go build

	// basics.IOOperation() //Working

	mathbasics.CountNumOfDigits(1234)
	basics.IOOperation()
	// var nimai = "Nimai"
	fmt.Println("hello")
	basics.PanicAndRecover()
	fmt.Println("test")
}
