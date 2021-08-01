package main

import "fmt"

func process(v int) string {
	return fmt.Sprintf("proc %d", v)
}

func processChannel(ch chan int) []string {
	const conc = 10
	results := make(chan string, conc)

	for i := 0; i < conc; i++ {
		go func(i int) {
			fmt.Printf("before receiving on goroutine %d\n", i)
			v := <-ch
			fmt.Printf("after receiving value %d on goroutine %d\n", v, i)
			results <- process(v)
		}(i)
	}

	var out []string
	for i := 0; i < conc; i++ {
		fmt.Printf("before appending on loop %d\n", i)
		v := <-results
		fmt.Printf("after reading value %s on loop %d\n", v, i)
		out = append(out, v)
	}

	return out
}

func main() {
	ch := make(chan int)

	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("sending %d\n", i)
			ch <- i
		}(i)
	}
	out := processChannel(ch)
	fmt.Println(out)
}
