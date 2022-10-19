package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/* Arquivos

	 */
	f, err := os.Create("file.txt")

	if err != nil {
		panic(err)
	}

	// write

	// length, err := f.WriteString("Hello, World!")
	length, err := f.Write([]byte("Writing bytes to the file"))

	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com successo! Tamanho: %d bytes\n", length)
	f.Close()

	// read

	file, err := os.ReadFile("file.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(file))

	// read line by line oppening the file

	file2, err := os.Open("file.txt")

	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file2)
	buffer := make([]byte, 2)

	for {
		n, err := reader.Read(buffer)

		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("file.txt")
	if err != nil {
		panic(err)
	}
}
