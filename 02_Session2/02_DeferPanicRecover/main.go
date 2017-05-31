package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Before poorFunc")
	if err := poorFunc(0); err != nil {
		fmt.Printf("Error: %s\n", err)
	}
	fmt.Println("After poorFunc")
}

func poorFunc(i int) error {
	if i == 0 {
		return errors.New("Cannot divide to 0!")
	}

	if 5/i == 0 {
		fmt.Println("Impossible")
	} else {
		fmt.Println("Possible")
	}

	return nil
}
