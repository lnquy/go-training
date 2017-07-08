package models

import "math"

type Circle struct {
	R float64 // Length of radius
}

func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

func (c *Circle) Perimeter() float64 {
	return 2.0 * math.Pi * c.R
}
