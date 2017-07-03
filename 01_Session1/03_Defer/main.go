package main

import (
	"fmt"
)

func main() {
	x := 0
	fmt.Printf("Original X: %v\n", x)
	defer fmt.Printf("X in defer: %v\n", x)

	x += 1
	fmt.Printf("X after add 1: %x\n", x)
	return
}

