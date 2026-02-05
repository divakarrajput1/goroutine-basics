package main

import (
	"fmt"
	"sync"
)

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 6; i++ {
		fmt.Println("Hello from goroutine")
		//time.Sleep(500 * time.Millisecond)
	}
}
func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go task1(&wg)

	wg.Wait()
}

// The main function runs line by line in order, just like normal code.
// First, it creates a WaitGroup, then it increases its counter to 1.
// Next, when it sees go task1(&wg), Go starts task1 as a separate goroutine and
// immediately moves on without waiting. After that, main reaches wg.Wait() and
// stops there, waiting. While main is waiting, the goroutine runs task1.
// When task1 finishes, it calls wg.Done(), the counter becomes 0, and wg.Wait() unblocks.
// Then main continues and the program exits cleanly
