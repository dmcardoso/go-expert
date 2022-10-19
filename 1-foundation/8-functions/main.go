package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(50, 10)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)
}

// func sum(a, b int) (int, error) { <- Função com parâmetros do mesmo tipo podem ser resumidas
func sum(a int, b int) (int, error) {
	sum := a + b
	if sum >= 50 {
		return sum, errors.New("A soma é maior que 50")
	}
	return sum, nil
}
