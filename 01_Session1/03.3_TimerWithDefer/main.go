package main

import (
	"time"
	"fmt"
)

func main() {
	longExecution1()
	longExecution2()
}

func longExecution1() {
	begin := time.Now()

	for i := 0; i < 10; i++ {
		time.Sleep(300 * time.Millisecond)
	}

	fmt.Printf("longExecution1 runs in: %v\n", time.Now().Sub(begin))
}

// ---------------------------------------------

func longExecution2() {
	defer timer("longExecution2", time.Now())

	for i := 0; i < 10; i++ {
		time.Sleep(300 * time.Millisecond)
	}
}

func timer(name string, t time.Time) {
	fmt.Printf("%s runs in: %v\n", name, time.Now().Sub(t))
}
