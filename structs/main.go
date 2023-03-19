package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	dan := person{
		firstName: "Dan",
		lastName:  "Henderson",
		contactInfo: contactInfo{
			email: "h@enderson.com",
			zip:   13704,
		},
	}

	dan.updateName("Daniel")
	dan.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
