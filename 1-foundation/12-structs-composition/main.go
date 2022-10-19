package main

import "fmt"

type Address struct {
	Street string
	Number int
	City   string
	State  string
}

type Client struct {
	Name    string
	Age     int
	Active  bool
	Address         // composition
	Add     Address // new property
}

func main() {
	// Structs composition

	daniel := Client{
		Name:   "Daniel",
		Age:    30,
		Active: true,
	}

	daniel.Active = false
	daniel.City = "Itapuranga"

	fmt.Println(daniel.Add.City)
	fmt.Println(daniel.Address.City)
	fmt.Println(daniel.City)
}
