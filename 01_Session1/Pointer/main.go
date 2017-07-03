package main

import (
	"fmt"
	"unsafe"
)

type XXX struct {
	Name string
}

func main() {

	var i int32 = 0

	fmt.Printf("Type: %T - Value: %v\n", i, unsafe.Sizeof(i))

	fmt.Printf("Type: %T - Value: %v\n", &i, unsafe.Sizeof(&i))

	//x := &i
	////	Name: "XXX",
	////}
	//
	//fmt.Printf("Type: %T - Value: %v\n", x, x)
	//fmt.Printf("Type: %T - Value: %v\n", *x, *x)
	//
	//var x [4]int = [4]int{0, 1, 2, 3}
	//
	//b := &x
	//
	//fmt.Printf("Type: %v - Value: %v", &b[0], &x)
}
