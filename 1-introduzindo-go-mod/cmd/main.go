package main

import (
	"fmt"
	"github.com/jonilsonds9/goexpert-modulo-5-packaging/1-introduzindo-go-mod/math"
)

func main() {
	m := math.Math{A: 1, B: 2}
	fmt.Println(m.Add())
	// fmt.Println("Hello, World!")
}