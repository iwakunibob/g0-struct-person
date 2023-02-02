package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // contact   contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	alex := person{"Alex", "Anderson", contactInfo{"AA@ah.com", 98765}} // No fields use order
	bob := person{
		lastName:  "Laurie",
		firstName: "Robert",
	}
	fmt.Println(alex)
	fmt.Println(bob)
	var deb person
	fmt.Println(deb)
	// %v+ prints field names and values
	fmt.Printf("%+v \n", deb)
	deb.firstName = "Debra"
	deb.lastName = "Novotny"
	fmt.Println(deb)
	deb.printFields()
	jim := person{
		firstName: "James",
		lastName:  "Laurie",
		contactInfo: contactInfo{ // contact contactInfo
			email:   "james@hotmail.com",
			zipCode: 95434,
		},
	}
	jim.print()
	ptrJim := &jim
	ptrJim.updateFName("Jimmy")
	jim.print()
}
func (p person) print() {
	fmt.Printf("%v \n", p)
}
func (p person) printFields() {
	fmt.Printf("%+v \n", p)
}
func (ptrPerson *person) updateFName(newFName string) {
	(*ptrPerson).firstName = newFName
}
