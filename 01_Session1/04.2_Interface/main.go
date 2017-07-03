package main

import (
	"gitlab.com/lnquy/go-training/01_Session1/04.2_Interface/models"
	"fmt"
)

type Sayer interface {
	Say() string
}

func main() {
	d := models.Dog{"Gogo"}
	c := models.Cat{"Mewmew"}

	LetAnimalSaySomething(d)
	LetAnimalSaySomething(c)
}

func LetAnimalSaySomething(a Sayer) {
	fmt.Println(a.Say())
}
