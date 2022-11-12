package main

import (
	"fmt"
	domain "test/domain"
)

func main() {

	address := domain.Address{Street: "123 SomeStreet", City: "MyCity", State: "California", Zip: 93213}
	person := domain.Person{FirstName: "Vlad", LastName: "Georgescu", Address: address}

	fmt.Printf(person.DisplayPerson())
}
