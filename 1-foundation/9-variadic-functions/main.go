package main

import (
	"fmt"
)

func main() {
	value := sum(50, 10, 30, 25, 55, 88, 5, 5, 6, 5, 854, 86, 1, 68, 81, 8, 1, 86, 168, 1, 8, 4168, 41, 8, 168, 16, 81, 65, 18, 41, 684, 68, 48)

	fmt.Println(value)
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}
	return total
}
