package main

import "fmt"

func main() {
	var m map[string]int = make(map[string]int)

	m["tma"] = 1
	m["fpt"] = 2
	m["aaa"] = 3

	if v, ok := m["xxx"]; ok {
		fmt.Printf("Ok: %v", v)
	} else {
		fmt.Printf("!Ok: %v", v)
	}

	//for k, v := range m {
	//	fmt.Printf("%s - %v", k, v)
	//}
}

