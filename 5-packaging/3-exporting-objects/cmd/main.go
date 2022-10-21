package main

import (
	"fmt"

	"github.com/dmcardoso/go-expert/5-packaging/3-exporting-objects/math"
)

func main() {
	m := math.NewMath(5, 6)

	fmt.Println(m.Add())
}
