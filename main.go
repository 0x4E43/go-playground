package main

import (
	"fmt"

	"github.com/0x4E43/go-playground/basics"
)

func main() {
	// Building aplication
	// go build
	// Extended: GOOS="linux" go build

	// basics.IOOperation() //Working
	// fmt.Println("HEllo")
	// mathbasics.CountNumOfDigits(1234)
	// basics.IOOperation()
	// var nimai = "Nimai"
	// fmt.Println("hello")
	// basics.PanicAndRecover()
	// fmt.Println("test")

	// basics.NeoVimTestHands()
	basics.StructureFunc()

	s := basics.Student{
		Name: "Hello",
		Age:  24,
	}
	name := s.AddName("World")
	fmt.Println(name)
}
