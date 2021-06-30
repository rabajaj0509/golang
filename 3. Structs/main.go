package main

import "fmt"

type contactInfo struct {
	address string
	pinCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Rahul",
		lastName:  "Bajaj",
		contactInfo: contactInfo{
			address: "Pune, Maharashtra",
			pinCode: 411013,
		},
	}
	jim.updateName("Shruti")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
