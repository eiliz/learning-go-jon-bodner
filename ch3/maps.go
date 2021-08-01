package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// intSet := map[int]bool{}
	// vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}

	// for _, v := range vals {
	// 	intSet[v] = true
	// }

	// fmt.Println(len(vals), len(intSet))
	// fmt.Println(intSet[5])
	// fmt.Println(intSet[500])

	// if intSet[100] {
	// 	fmt.Println("100 is in the set")
	// }

	// intSet := map[int]struct{}{}
	// vals := []int{5, 10, 2, 5, 8, 7, 3, 9, 1, 2, 10}
	// for _, v := range vals {
	// 	intSet[v] = struct{}{}
	// }

	// if _, ok := intSet[5]; ok {
	// 	fmt.Println("5 is in the set")
	// }

	// fmt.Println(intSet)

	var a HasEmptyStructAtStart
	var b HasEmptyStructAtEnd

	a.addr()                      // 0xc0000140a0 0xc0000140a0 0xc0000140a0
	fmt.Println(unsafe.Sizeof(a)) // 8

	b.addr()                      // 0xc0000140b0 0xc0000140b8 0xc0000140b0
	fmt.Println(unsafe.Sizeof(b)) // 16

	// Struct type instances are comparable with == if all the struct's
	// fields are comparable types like strings, it doesn't work for slice or map
	// fields.

	// If they have the same fields - name, types and order, you can convert from
	// different named types of structs.

	// You can also compare and assign values from one type to another between a
	// named struct and an annonymous struct if they have the same fields - name,
	// types, order and do that without a conversion like in the case of named
	// structs.

	type firstPerson struct {
		name string
		age  int
	}

	f := firstPerson{
		name: "Bob",
		age:  50,
	}

	var g struct {
		name string
		age  int
	}

	h := struct {
		name string
		age  int
	}{
		name: "Bob",
		age:  50,
	}

	// compiles -- can use = and == between identical named and anonymous structs
	g = f
	fmt.Println(f == g)
	f = h
	fmt.Println(f == h)

	ages := make(map[int][]string, 10)
	fmt.Println(len(ages))
}

// This will have a width of 8 bytes
type HasEmptyStructAtStart struct {
	Empty struct{}
	Name  int
}

// This will have a width of 16 bytes
type HasEmptyStructAtEnd struct {
	Name  int
	Empty struct{}
}

// The difference in size is because the rules say for every struct field
// you could take a pointer to it. And every pointer is registered by the GC.
// So even if a struct{} type field has no width it still needs to be
// addressable.
// If you put it at the start of a struct it can use the start address of the
// struct. If it's at the end it would mean that you might address some other
// random memmory and cause errors so 1 byte is allocated for the empty struct
// and that's why we get an extra 8 bytes added to the total width - the arch is
// 64 and the max alignment in the struct comes from the alignment of the field
// with the largest alignment, the int in this case.
// So it's 8 bytes and you get 7 extra bytes of padding after the struct{}'s 1
// bytes to get the right alignment to 8.

func (s *HasEmptyStructAtStart) addr() { fmt.Printf("%p %p %p\n", s, &s.Empty, &s.Name) }

func (s *HasEmptyStructAtEnd) addr() { fmt.Printf("%p %p %p\n", s, &s.Empty, &s.Name) }
