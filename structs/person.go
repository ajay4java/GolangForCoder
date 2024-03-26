package structs

type Person struct {
	firstName string
	lastName  string
	age       int
}

func NewPerson(firstName, lastName string, age int) Person {

	return Person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
	}

}

func (p Person) Walk() string {

	return p.firstName + p.lastName + " loves to walk"

}

func (p *Person) SetFirstName(f string) {
	p.firstName = f
}
func (p Person) SetLastName(l string) {
	p.lastName = l
}

func (p Person) SetAge(a int) {
	p.age = a
}
