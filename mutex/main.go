package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int
var mu sync.Mutex

func increment() {
	mu.Lock()
	defer mu.Unlock()

	counter++

	fmt.Println("Counter:", counter)

}

func main() {
	for i := 0; i < 5; i++ {
		go increment()
	}

	time.Sleep(time.Second * 1)
}
