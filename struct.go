package main

import "fmt"

type Nimai struct {
	name string
	age  int
}

func main11() {
	fmt.Println("Hello Struct test")

	// nimai := Nimai{"Nimai", 24}

	var test []Nimai
	test = append(test, Nimai{"Nimai", 24})
	test = append(test, Nimai{"Charan", 24})
	test = append(test, Nimai{"Maikap", 24})

	fmt.Println("Struct List length", len(test))
	// printStruct(nimai)
	// printStructSlice(test)
	per := CreatePerson()
	fmt.Println("Before Edit (WITHOUT POINTER): ", per)
	editPerson(per)
	fmt.Println("After Edit (WITHOUT POINTER): ", per)

	perPnt := CreatePerson()
	fmt.Println("Before Edit (WITH POINTER): ", perPnt)

	editPersonPnt(&perPnt)

	fmt.Println("After Edit(WITH POINTER)L ", perPnt)
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

func editPerson(person Person) {
	person.age = 45
	// return person
}

func editPersonPnt(per *Person) {
	per.age = 45
}
