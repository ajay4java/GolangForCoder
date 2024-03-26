package polymorphism

import "fmt"

type Shape interface {
	Render()
}

type Circle struct {
}

func (c Circle) Render() {
	fmt.Println("Circle")

}

type Square struct {
}

func (s Square) Render() {
	fmt.Println("Square")

}
