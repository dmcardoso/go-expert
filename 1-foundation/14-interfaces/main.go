package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name   string
	Age    int
	Active bool
	Address
}

func (c Client) SetDisabled() {
	c.Active = false
	fmt.Printf("O cliente %s está com status: %v\n", c.Name, c.Active)
}

type Company struct {
	Name string
}

func (c Company) SetDisabled() {}

/*
Interfaces are implemented automatically

Interfaces can't have properties, only methods
*/
type Person interface {
	SetDisabled()
}

func Disable(person Person) {
	person.SetDisabled()
}

func main() {
	// Interfaces

	daniel := Client{
		Name:   "Daniel",
		Age:    30,
		Active: true,
	}

	company := Company{
		Name: "iFood",
	}

	Disable(daniel)
	Disable(company)
}
