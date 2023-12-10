package main

import "fmt"

type Nimai struct {
	name string
	age  int
}

func main() {
	fmt.Println("Hello Struct test")

	// nimai := Nimai{"Nimai", 24}

	var test []Nimai
	test = append(test, Nimai{"Nimai", 24})
	test = append(test, Nimai{"Charan", 24})
	test = append(test, Nimai{"Maikap", 24})

	fmt.Println("Struct List length", len(test))

	printStruct(test)
}

func printStruct(structList []Nimai) {

	for i := 0; i < len(structList); i++ {
		fmt.Println("Name: ", structList[i].name, " Age: ", structList[i].age)
	}

}
