package basics

import "fmt"

type Student struct{
  Name string
  Age int8
}


func StructureFunc(){

  str := Student{
    Name: "Nimai",
    Age: 24,
  }

  fmt.Printf("Name: %s | Age : %d ",str.Name, str.Age)

}

func (s *Student) AddName( name string)(string){
  s.Name = name
  return s.Name
}
