package main

const a = "Hello World"

var (
	b bool    = true
	c int     = 10
	d string  = "Daniel"
	e float64 = 1.2
)

func main() {
	b = true

	// shorthand - declaração por inferência
	a := "x"

	// mutação de valor
	a = "fd"

	println(a)
}

func x() {}
