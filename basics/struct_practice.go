package basics

import "fmt"

type Student struct{
  name string
  age int8
}


func StructureFunc(){

  str := Student{
    name: "Nimai",
    age: 24,
  }

  fmt.Printf("Name: %s | Age : %d ",str.name, str.age)

}
