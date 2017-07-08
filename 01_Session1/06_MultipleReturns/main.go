package main

import "fmt"

func main() {
	a, b := 0, 10
	fmt.Printf("Original: A=%d - B=%d\n", a, b)

	a, b = reverse(a, b)
	fmt.Printf("Reversed: A=%d - B=%d\n", a, b)

	//a, b = b, a
}

func reverse(a, b int) (int, int) {
	return b, a
}
