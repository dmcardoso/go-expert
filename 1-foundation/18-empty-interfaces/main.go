package main

import "fmt"

type x interface{}

func main() {
	/** Empty interfaces
	 */

	var x interface{} = 10
	var y interface{} = "Hello, World!"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e p valor é %v\n", t, t)
}
