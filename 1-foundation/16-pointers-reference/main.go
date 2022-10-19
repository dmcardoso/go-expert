package main

func sum(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	/** Pointers

	Usage
	*/

	myVar1 := 10
	myVar2 := 20

	println(sum(&myVar1, &myVar2), myVar1, myVar2)

}
