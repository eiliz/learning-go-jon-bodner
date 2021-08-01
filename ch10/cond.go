package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	queue := make([]interface{}, 0, 10)
	c := sync.NewCond(&sync.Mutex{})

	removeFromQueue := func(delay time.Duration, i int) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Printf("%d removed from the queue\n", i)
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 5 {
			c.Wait()
		}
		fmt.Printf("%d adding to queue\n", i)
		queue = append(queue, struct{}{})
		go removeFromQueue(1*time.Second, i)
		c.L.Unlock()
	}
}
