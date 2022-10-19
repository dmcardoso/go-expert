package main

func main() {
	/* For */
	for i := 0; i < 10; i++ {
		println(i)
	}

	numbers := []string{"um", "dois", "três"}

	println("\n=============\n")

	for k, v := range numbers {
		println(k, v)
	}

	println("\n=============\n")

	for k := range numbers {
		println(k)
	}

	println("\n=============\n")

	for _, v := range numbers {
		println(v)
	}

	println("\n=============\n")

	i := 0
	for i < 10 {
		println(i)
		i++
	}

	println("\n=============\n")

	for {
		println("Hello, World!")
	}

}
