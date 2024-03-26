package composition

type engine struct {
	hp int
}

func (e engine) HP() int {
	return e.hp
}

type wheel struct {
	wheelDimensions int
}

func (w wheel) WheelDimensions() int {
	return w.wheelDimensions
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
