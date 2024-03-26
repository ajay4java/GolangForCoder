package composition

type engine struct {
	hp int
}
type wheel struct {
	wheelDimensions int
}
type Car struct {
	CarName string
	Engine  engine
	Wheel   wheel
}

func NewCar(carName string, hp, wd int) Car {
	return Car{
		CarName: carName,
		Engine: engine{
			hp: hp,
		},
		Wheel: wheel{
			wheelDimensions: wd,
		},
	}
}
