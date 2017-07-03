package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Duration(i*100) * time.Millisecond)
	}
}

func main() {
	f("direct")

	go f("goroutine")
	go f("qip")

	//go func(msg string) {
	//	fmt.Println(msg)
	//}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
