package main

import "fmt"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Daniel"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de F é %T\n", f)
	fmt.Printf("O valor de F é %v", f)
}
