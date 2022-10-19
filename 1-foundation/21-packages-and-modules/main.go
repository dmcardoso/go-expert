package main

import (
	"fmt"

	"go-expert/maths"

	"github.com/google/uuid"
)

func main() {
	/** Packages and modules
	Functions are only exported when starts with a capital letter
	*/

	sum := maths.Sum(10, 20)

	car := maths.Car{Brand: "Fiat"}

	fmt.Println("Resultado: ", sum)
	fmt.Println("maths.A: ", maths.A)
	fmt.Println("car.Brand: ", car.Brand)
	fmt.Println("car: ", car)

	fmt.Println(uuid.New())
}
