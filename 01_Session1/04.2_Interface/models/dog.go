package models

import "fmt"

type Dog struct {
	Name string
}

func (d Dog) Say() string {
	return fmt.Sprintf("[%s] I'm a goodboy =))", d.Name)
}
