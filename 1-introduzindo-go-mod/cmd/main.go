package main

import (
	"fmt"
	"github.com/jonilsonds9/goexpert-modulo-5-packaging/1-introduzindo-go-mod/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.C)
	// fmt.Println(m.Add())
	// fmt.Println(math.X)
}