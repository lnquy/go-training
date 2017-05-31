package main

import "fmt"

func main() {
	var arr [5]string = [5]string{"One", "Two", "Three", "Four", "Five"}
	for  i, v := range arr {
		fmt.Printf("Index: %v => Value: %v\n", i, v)
	}

	fmt.Println("===== MAP =====")
	var m map[int]string = map[int]string{10:"Bill Gates", 100:"Dijkstra", 1000:"Alan Turing"}
	for k, v := range m {
		fmt.Printf("Key: %v => Value: %v\n", k, v)
	}
}
