package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	ContactInfo
}

type ContactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := Person{
		firstName: "Jim",
		lastName:  "Party",
		ContactInfo: ContactInfo{
			email:   "jim@gmail.com",
			zipCode: 94000,
		},
	}

	jimPointer := &jim
	jimPointer.updateName("jimmy")

	jim.print()

	fmt.Printf("Jim Pointer %v", jimPointer)

	// &variable - Give me the memory address of the value this variable pointing at
	// *pointer - Give me the value this memory address is pointing at
	// Turn address into value with *address
	// Turn value into address with &value
}

func (pointerToPerson *Person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p Person) print() {
	fmt.Printf("%+v \n", p)
}
