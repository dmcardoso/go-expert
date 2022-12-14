package main

func main() {
	ch := make(chan string, 2)

	ch <- "Hello"
	ch <- "Hello"

	println(<-ch)
	println(<-ch)
}
