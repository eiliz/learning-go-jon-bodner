package main

import (
	"fmt"
)

func main() {
	// Appending is done based on the length, not the capacity so if a slice has
	// the default value and a length of 5, it's going to be populated with 5 0's
	// and appending to it will start after the 5 0's.

	// Makes a new slice with a capacity of 2 - 0 = 2, the 3rd param value - the
	// offset.
	// Because the capacity is smaller than that of the intial slice, x and y
	// cannot share the same underlying array and therefore the subslice y is
	// disconnected from x and changes to it won't affect x.
	// x := make([]int, 0, 5)
	// x = append(x, 1, 2, 3, 4)
	// y := x[:2:2]
	// z := x[2:4:4]
	// fmt.Println(cap(x), cap(y), cap(z))
	// y = append(y, 30, 40, 50)
	// x = append(x, 60)
	// z = append(z, 70)
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)
	// fmt.Println("z:", z)

	// x := [4]int{5, 6, 7, 8}
	// y := x[:2]
	// y = append(y, 30, 40, 50)
	// x[0] = 10
	// fmt.Println("x:", x)
	// fmt.Println("y:", y)

	// var s string = "Hello, 😀"

	// for _, c := range s {
	// 	fmt.Printf("%s", string(c))
	// }

	// for i := 0; i < len(s); i++ {
	// 	fmt.Printf("%s", string(s[i]))
	// }
	var a byte = 'a'
	fmt.Printf("%T", a)
}
