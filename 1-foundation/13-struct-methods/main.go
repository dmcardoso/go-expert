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

// Might use a pointer to change the original struct
func (c Client) SetDisabled() {
	c.Active = false
	fmt.Printf("O cliente %s está com status: %v\n", c.Name, c.Active)
}

func main() {
	// Structs methods

	daniel := Client{
		Name:   "Daniel",
		Age:    30,
		Active: true,
	}

	daniel.City = "Itapuranga"
	daniel.SetDisabled()

	fmt.Println(daniel.Active)
}
