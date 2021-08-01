package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func main() {
	// Whenever a new map variable is created Go will generate a random number to
	// be used for hashing the keys of that specific map. This means that 2
	// different map variables with the same keys won't have the same hashes. This
	// prevents the Hash DoS attack that would send some data to the server that
	// would make all the keys hash to the same bucket.

	// The for range loop is also random to help with the same problem.
	// When sending a map to a fmt function it will print its keys in ascending
	// order to help with debugging.
	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }

	// for i := 0; i < 3; i++ {
	// 	fmt.Println("Loop", i)
	// 	for k, v := range m {
	// 		fmt.Println(k, v)
	// 	}
	// }

	// fmt.Println(m)

	// The for-range loop iterates over your compound type, it copies the value from the compound type to the value variable
	// evenVals := []int{2, 4, 6, 8, 10, 12}
	// for _, v := range evenVals {
	// 	v *= 2
	// }
	// fmt.Println(evenVals)

	// evenVals2 := []int{2, 4, 6, 8, 10, 12}
	// for i, v := range evenVals2 {
	// 	evenVals2[i] = 2 * v
	// }
	// fmt.Println(evenVals2)

	// type Test struct {
	// 	a int
	// }

	// src := []Test{{1}, {2}, {3}, {4}}

	// for _, v := range src {
	// 	v.a *= 2
	// 	fmt.Print(v)
	// }

	// fmt.Println()

	// for _, v := range src {
	// 	fmt.Print(v)
	// }
	// // {2}{4}{6}{8}
	// // {1}{2}{3}{4}

	// src := []*Test{{1}, {2}, {3}, {4}}

	// for _, v := range src {
	// 	v.a *= 2
	// 	fmt.Print(v)
	// }

	// fmt.Println()

	// for _, v := range src {
	// 	fmt.Print(v)
	// }
	// // &{2}&{4}&{6}&{8}
	// // &{2}&{4}&{6}&{8}

	// Regular switches will use the equality comparison for the check.
	words := []string{"a", "cow", "smile", "gopher",
		"octopus", "anthropologist"}
	for _, word := range words {
		switch size := len(word); size {
		case 1, 2, 3, 4:
			fmt.Println(word, "is a short word!")
		case 5:
			fmt.Println(word, "is exactly the right length:", size)
		case 6, 7, 8, 9:
		default:
			fmt.Println(word, "is a long word!")
		}
	}

	// Blank switches allow you to perform any boolean comparison inside the cases
	// themselves.
	otherWords := []string{"hi", "salutations", "hello"}
	for _, word := range otherWords {
		switch wordLen := len(word); {
		case wordLen < 5:
			fmt.Println(word, "is a short word!")
		case wordLen > 10:
			fmt.Println(word, "is a long word!")
		default:
			fmt.Println(word, "is exactly the right length.")
		}
	}

	// Also a blank switch that includes a var declaration. The var is scoped to
	// the entire switch block.
	switch n := rand.Intn(10); {
	case n == 0:
		fmt.Println("That's too low")
	case n > 5:
		fmt.Println("That's too big:", n)
	default:
		fmt.Println("That's a good number:", n)
	}

loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println(i, "is even")
		case i%3 == 0:
			fmt.Println(i, "is divisible by 3 but not 2")
		case i%7 == 0:
			fmt.Println("exit the loop!")
			break loop
		default:
			fmt.Println(i, "is boring")
		}
	}

	test()()

	type Person struct {
		FirstName string
		LastName  string
		Age       int
	}

	people := []Person{
		{"Pat", "Patterson", 37},
		{"Tracy", "Bobbert", 23},
		{"Fred", "Fredson", 18},
	}

	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].LastName < people[j].LastName
	})

	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println(people)
}

func test() func() {
	hidden := 7

	return func() {
		fmt.Println(hidden)
	}
}
