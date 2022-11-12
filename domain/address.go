package domain

import "fmt"

type Address struct {
	Street string
	City   string
	State  string
	Zip    int
}

func (address Address) DisplayAddress() string {
	return fmt.Sprintf(">>> %v %v %v %d.\n", address.Street, address.City, address.State, address.Zip)
}
