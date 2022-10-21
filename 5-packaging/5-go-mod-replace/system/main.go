package main

import "github.com/dmcardoso/go-expert/5-packaging/5-go-mod-replace/math"

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())
}
