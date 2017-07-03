package models

import "fmt"

type Cat struct {
	Name string
}

func (c Cat) Say() string {
	return fmt.Sprintf("[%s] I'm a fluffy cat ~.~", c.Name)
}
