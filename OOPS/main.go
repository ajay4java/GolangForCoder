package main

import (
	"GolangForCode/composition"
	"fmt"
)

func main() {
	c := composition.NewCar("testCar", 600, 123)
	fmt.Println(c)
	fmt.Println(c.Engine.HP())
	fmt.Println(c.Wheel.WheelDimensions())

}
