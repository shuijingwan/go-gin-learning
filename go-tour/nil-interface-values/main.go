package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S string
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
