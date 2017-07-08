/*Write an application that defines 4 structs for square, rectangle, triangle and circle shape.
	Each struct type has 2 behaviors: area (calculate area) and perimeter (calculate perimeter) that
	implemented from the Calculator interface.
Write an function with signature of: func measure(c Calculator) that accepts Calculator interface
	as the argument then print out the area and perimeter of the c input.

E.g:
type Calculator interface {
   Area() float64
   Perimeter() float64
}
*/

package main

import (
	"fmt"
	"gitlab.com/lnquy/go-training/03_Others/03_Homework_170703/Practice2/models"
	"reflect"
)

type Calculator interface {
	Area() float64
	Perimeter() float64
}

func main() {
	// Notice: The Area() and Perimeter() methods has been implemented for *models.Circle type, not models.Circle type.
	// The same for Rectangle, Square and Triangle.
	// Let try to remove the & symbols before models.Circle and re-run to see what happen
	c := &models.Circle{
		R: 3.55,
	}
	r := &models.Rectangle{
		L: 60.3,
		W: 34.2,
	}
	s := &models.Square{
		L: 3,
	}
	t := &models.Triangle{
		A: 3,
		B: 4,
		C: 5.0,
	}

	// Since both *Circle, *Rectangle, *Square and *Triangle already implemented the Area() and Perimeter() methods
	// from Calculator interface, it can be used as Calculator type
	calEverything(c)
	calEverything(r)
	calEverything(s)
	calEverything(t)
}

func calEverything(cal Calculator) {
	fmt.Printf("\nType of shape: %v\n", reflect.TypeOf(cal)) // Reflection - Get actual type of struct that implemented Calculator interface
	fmt.Printf("  Values: %v\n", cal)
	fmt.Printf("  Area: %.3f\n", cal.Area())
	fmt.Printf("  Primeter: %.3f\n", cal.Perimeter())
}
