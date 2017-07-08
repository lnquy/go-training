package models

type Rectangle struct {
	L float64 // Length of length side
	W float64 // Length of width side
}

func (r *Rectangle) Area() float64 {
	return r.L * r.W
}

func (r *Rectangle) Perimeter() float64 {
	return 2.0 * (r.L + r.W)
}
