package main

import (
	"io"
	"net/http"
)

func main() {
	/* Defer

	Called in the end of the function

	*/
	req, err := http.Get("https://google.com")

	if err != nil {
		panic(err)
	}

	defer req.Body.Close()

	res, err := io.ReadAll(req.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))

	println("First line")
	defer println("Second line")
	println("Third line")
}
