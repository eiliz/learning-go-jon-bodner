package main

import (
	"fmt"
)

func main() {
	// 0000 0000
	// 0000 0001
	// 0000 0010
	// 0000 0011
	// fmt.Println(popcount.PopCount(15096512))
	// fmt.Println(popcount.PopCountIn64Steps(15096512))
	// fmt.Println(1 & 2)

	a, b := 1, 2
	b, c := 3, 4

	if a := 7; a > 4 {
		fmt.Println("yup, locally scoped")
	}

	fmt.Println(a, b, c)
}
