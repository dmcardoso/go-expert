package main

import (
	"github.com/dmcardoso/go-expert/5-packaging/6-workspaces/math"
	"github.com/google/uuid"
)

func main() {
	m := math.NewMath(1, 2)

	println(uuid.New().String())
	println(m.Add())
}
