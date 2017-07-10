package main

import (
	"fmt"
	"log"
)

func main() {
	defer poorFunc()

	fmt.Println("Begin main")
	fmt.Println("End main")
}

func poorFunc() {
	log.Panic("Stupid panic inside defer")
}
