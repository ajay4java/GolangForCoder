package main

import "fmt"

func main() {

	const (
		name = "Ajay"
		age  = 45
	)
	fmt.Println("name=", name, "age=", age)
	fmt.Printf("name is of %T \n", name)
	fmt.Printf("name is of %T \n", age)

	var a float64 = 10
	var b int = 3
	var c int = int(a) + b
	fmt.Println(c)

}
