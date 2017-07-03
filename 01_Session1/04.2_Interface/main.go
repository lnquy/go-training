package main

import (
	"fmt"
)

type Sayer interface {
	Say() string
}

type FuncType func(...string)

func X1(s ...string) {
	fmt.Print(s)
}

func X2(s ...string) {
	fmt.Printf("XXX: %s", s)
}

func main() {
	//d := models.Dog{"Gogo"}
	//c := models.Cat{"Mewmew"}
	//
	//LetAnimalSaySomething(d)
	//LetAnimalSaySomething(c)

	X(X1)
	X(X2)
}

func X(f FuncType) {
	f("asdas", "asdas, 'q;weqwe")
}

func LetAnimalSaySomething(a Sayer) {
	fmt.Println(a.Say())
}
