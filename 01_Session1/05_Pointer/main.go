package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	user := retUser()
	fmt.Printf("Type: %T\nValue: %v", user, user)
}

func retUser() *User {
	return &User{
		Name: "lnquy",
		Age:  24,
	}
}
