package main

func main() {
	/** Pointers

	Memory -> Address -> Value

	*/

	a := 10

	var pointer *int = &a
	println(pointer)

	println(a)

	*pointer = 20
	println(a)

	b := &a
	println(b)
	println(*b)

	*b = 30
	println(a)
}
