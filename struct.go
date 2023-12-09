package main

import "fmt"

type Nimai struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello Struct test")

	nimai := Nimai{"Nimai", 24}

	var test []Nimai
	test[0] = Nimai{"Nimai", 24}
	test[1] = Nimai{"Charan", 24}
	test[3] = Nimai{"Maikap", 24}

	fmt.Println("Struct Values: ", nimai.name)
}
