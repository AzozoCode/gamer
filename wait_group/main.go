package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var cond = sync.NewCond(&mu)
var ready = false

func worker(id int) {
	mu.Lock()
	for !ready {
		cond.Wait()
	}
	mu.Unlock()
	fmt.Printf("Worker %d is running\n", id)
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			defer wg.Done()
			worker(id)
		}(i)
	}

	mu.Lock()
	ready = true
	cond.Broadcast() // signal to wake all the gorountines
	mu.Unlock()

	wg.Wait()
}
