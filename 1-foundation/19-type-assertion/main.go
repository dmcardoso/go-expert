package main

import "fmt"

func main() {
	/** Type assertion
	 */

	var myVar interface{} = "Daniel Moreira"

	println(myVar.(string))

	res, ok := myVar.(int)
	fmt.Printf("O valor de res é %v e o resultado de ok é %v\n", res, ok)

	res2 := myVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
