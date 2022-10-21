package main

import (
	"fmt"

	"github.com/dmcardoso/go-expert/5-packaging/2-accessing-packages/math"
)

func main() {
	m := math.Math{A: 5, B: 6}

	fmt.Println(m.Add())
}
