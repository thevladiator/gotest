package domain

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Address   Address
}

func (person Person) DisplayPerson() string {
	return fmt.Sprintf("Hi, %v %v!\n%v", person.FirstName, person.LastName, person.Address.DisplayAddress())
}
