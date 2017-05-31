package main

import "fmt"

func main() {
	fmt.Println("Before poorFunc")
	poorFunc(0)
	fmt.Println("After poorFunc")
}

func poorFunc(i int) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("Goroutine was panicked (%s)! But we can recover from it :))\n", err)
		} else {
			fmt.Println("Goroutine is ok!")
		}
	}()

	if 5 / i == 0 {
		fmt.Println("Impossible")
	} else {
		fmt.Println("Possible")
	}
}
