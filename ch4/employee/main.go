package main

import (
	"fmt"
	"unsafe"
)

type Cat struct {
	Type int32
	Name string
}

func main() {
	cat := Cat{Type: 1, Name: "Garfield"}
	fmt.Printf("%p\n%p\n", &cat.Type, &cat.Name)
	fmt.Println(unsafe.Sizeof(cat))
	fmt.Println(int(0xc0000a6018), int(0xc0000a6020))
}
