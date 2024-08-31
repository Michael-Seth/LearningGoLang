package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}
type secretAgent struct {
	person
	licenseToKill bool
}

type human interface {
	speak()
}

func main() {
	// fmt.Println("Hi There")

	secretAgent1 := secretAgent{
		person{
			"Michael",
			"Osas",
		},
		true,
	}
	// secretAgent1.speak()
	// secretAgent1.person.speak()
	saySomething(secretAgent1)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.firstName, sa.lastName, `Says, "shaken, not stirred"`)
}

func (p person) speak() {
	fmt.Println(p.firstName, p.lastName, `Says, "Good morning Michael"`)

}

func saySomething(h human) {
	h.speak()
}
