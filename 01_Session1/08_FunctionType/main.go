package main

import "fmt"

type Calculator struct {
        acc float64
}

type FuncType func(float64, float64) float64

func (c *Calculator) Do(ft FuncType, v float64) float64 {
        c.acc = ft(c.acc, v)
        return c.acc
}


func Add(a, b float64) float64 { return a + b }
func Sub(a, b float64) float64 { return a - b }
func Mul(a, b float64) float64 { return a * b }

func main() {
        var c Calculator
        fmt.Println(c.Do(Add, 5))       // 5
        fmt.Println(c.Do(Sub, 3))       // 2
        fmt.Println(c.Do(Mul, 8))       // 16
}