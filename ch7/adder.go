package main

import "fmt"

type Adder struct {
	start int
}

func (a Adder) AddTo(val int) int {
	return a.start + val
}

func main() {
	adder := Adder{start: 7}
	f1 := adder.AddTo
	fmt.Println(f1(10))
	adder.start = 12
	fmt.Println(adder)
	fmt.Println(f1(10))
}
