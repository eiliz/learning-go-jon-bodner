package main

import "fmt"

type IntTree struct {
	val         int
	left, right *IntTree
}

func (it *IntTree) Insert(v int) *IntTree {
	if it == nil {
		return &IntTree{val: v}
	}

	if v < it.val {
		it.left = it.left.Insert(v)
	} else if v > it.val {
		it.right = it.right.Insert(v)
	}

	return it
}

func (it *IntTree) Contains(v int) bool {
	switch {
	case it == nil:
		return false

	case v < it.val:
		return it.left.Contains(v)

	case v > it.val:
		return it.right.Contains(v)

	default:
		return true
	}
}

func main() {
	var it *IntTree
	it = it.Insert(5)
	it = it.Insert(3)
	it = it.Insert(10)
	it = it.Insert(2)
	fmt.Println(it.Contains(2))  // true
	fmt.Println(it.Contains(12)) // false
}
