package main

import "fmt"

type (
	Student struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	Person struct {
		Student  // Embedded / Anonymous struct

		Height int
	}
)

func (s *Student) Say() {
	fmt.Printf("I'm student %s", s.Name)
}

func (s *Student) ChangeName(name string) {
	s.Name = name
}

func main() {
	s := Student{
		Age:  12,
		Name: "AAA",
	}

	s.Say()
	s.ChangeName("xxx")
	s.Say()

	p := Person{
		Student: s,
		Height: 177,
	}

	p.Say()
}
