package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Mover interface {
	Move()
}

type Student struct {
	Person // Embedded struct

	Class string
	School string

	Mover // Embedded interface
}

func (*Student) Move() {
	fmt.Println("Student moves...")
}


func main() {
	s := &Student{
		Person: Person{
			Name: "lnquy",
			Age: 24,
		},
		Class: "11CN111",
		School: "LHU",
	}

	fmt.Printf("Name: %s\nAge: %d\nClass: %s\nSchool: %s\n", s.Person.Name, s.Age, s.Class, s.School)
	s.Move()
}