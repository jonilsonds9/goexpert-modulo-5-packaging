package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jonilsonds9/goexpert-modulo-5-packaging/4-usando-workspaces/math"
)

func main() {
	m := math.NewMath(1, 2)
	fmt.Println(m.Add())
	fmt.Println(uuid.NewString())
}