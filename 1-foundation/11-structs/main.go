package main

import "fmt"

type Client struct {
	Name   string
	Age    int
	Active bool
}

func main() {
	// Structs

	daniel := Client{
		Name:   "Daniel",
		Age:    30,
		Active: true,
	}

	daniel.Active = false

	fmt.Println(daniel.Active)
}
