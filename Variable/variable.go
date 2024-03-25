package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b float64 = 10, 20
	var (
		name1 = "Ruchita"
		age1  = 50
	)
	var age int = 29
	name := "Ajay"

	fmt.Println(age)
	fmt.Println(name)

	fmt.Println("Name = ", name, " Age = ", age)
	fmt.Println("Name = ", name1, " Age = ", age1)
	fmt.Println("a =", a, "b =", b)

	fmt.Println(math.Min(a, b))
}
