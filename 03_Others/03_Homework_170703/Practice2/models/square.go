package models

import "math"

type Square struct {
	L float64 // Length of a side
}

func (s *Square) Area() float64 {
	return math.Pow(s.L, 2.0)
}

func (s *Square) Perimeter() float64 {
	return 4.0 * s.L
}
