package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	jim := person{
		firstName: "James",
		lastName:  "Laurie",
	}
	jim.print()
	// ptrJim := &jim  // Go is smart not needed
	// ptrJim.updateFName("Jimmy")
	jim.updateFName("Jimmy")
	jim.print()
}
func (p person) print() {
	fmt.Printf("%v \n", p)
}
func (ptrPerson *person) updateFName(newFName string) {
	(*ptrPerson).firstName = newFName
}
