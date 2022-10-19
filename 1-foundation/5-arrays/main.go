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
	// Arrays in go has a fixed length
	var myArray [4]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30
	myArray[3] = 40

	fmt.Println(len(myArray))
	fmt.Println(myArray[len(myArray)-1])
	fmt.Println(myArray[1])

	for index, value := range myArray {
		fmt.Printf("O valor do índice é %d e o valor é %d\n", index, value)
	}
}
