package models

import "math"

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t *Triangle) Area() float64 {
	p := t.Perimeter() / 2.0
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}
