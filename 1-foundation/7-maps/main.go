package main

import "fmt"

func main() {
	// Maps
	salaries := map[string]int{"Daniel": 1000, "João": 2000, "Maria": 3000}

	fmt.Println(salaries["Daniel"])

	delete(salaries, "Daniel")
	salaries["Dan"] = 5000
	fmt.Println(salaries["Dan"])

	salaryMaker := make(map[string]int)
	salaryMaker1 := map[string]int{}

	salaryMaker["Dan"] = 5000
	salaryMaker1["Dan"] = 5000

	for name, salary := range salaries {
		fmt.Printf("O salário de %s é %d\n", name, salary)
	}
	for _, salary := range salaries {
		fmt.Printf("O salário é %d\n", salary)
	}
}
