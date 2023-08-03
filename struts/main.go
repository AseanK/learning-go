package main

import "fmt"

type contact struct {
	email string
	phone int
}

type person struct {
	firstName string
	lastName  string
	age       int
	contact
}

func main() {
	// Rely on the order of the struct
	john := person{
		"John",
		"Doe",
		21,
		contact{
			"email@email.com",
			123123123,
		},
	}

	// Give keyword to specify
	alex := person{
		lastName:  "Doe",
		age:       32,
		firstName: "Alex",
		contact: contact{
			email: "admin@admin.com",
			phone: 321321321,
		},
	}

	// Assign the strut to the value first
	// If not assigned, Zero Value is assigned
	// str: "", int:0, float:0, bool:false
	var jim person
	// fmt.Printf("%+v", jim)

	jim.firstName = "Jim"
	jim.lastName = "Gray"
	jim.age = 53
	jim.contact.email = "jim@gary.com"
	jim.contact.phone = 111222333

	alex.print()
	john.print()

	// &var : Memory address. Pointer to the RAM address
	jim.updateName("Jym")
	jim.print()
	fmt.Printf("%p", &jim)
}

// *: Takes RAM address and gives the value
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

// Strut receiver function
func (p person) print() {
	fmt.Printf("%+v\n", p)
}
