package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstname string
	lastname  string
	contact   contactInfo
}

func main() {
	jim := person{
		firstname: "Jim",
		lastname:  "Nashburn",
		contact: contactInfo{
			email:   "nashburn_jim@gmail.com",
			zipcode: 34343,
		},
	}
	jim.print()
	jimPointer := &jim
	jimPointer.updateName("Jimmy")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(name string) {
	(*p).firstname = name
}
