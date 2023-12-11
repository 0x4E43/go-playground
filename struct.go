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
	// printStruct(nimai)
	// printStructSlice(test)

	fmt.Println(CreatePerson().name)
}

func printStruct(nimai Nimai) {
	fmt.Println("Name: ", nimai.name, " Age: ", nimai.age)
}

func printStructSlice(structList []Nimai) {

	for i := 0; i < len(structList); i++ {
		fmt.Println("Name: ", structList[i].name, " Age: ", structList[i].age)
	}

}

type Person struct {
	name   string
	age    int
	mobile string
}

func CreatePerson() Person {
	per := Person{"Nimai Charan Maikap", 24, ""}
	return per
}
