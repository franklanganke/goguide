package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	lastname  string
	contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Januar",
		contactInfo: contactInfo{
			email:   "jim@localhost",
			zipCode: 22337},
	}

	jim.updateName("jimmy")
	jim.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
