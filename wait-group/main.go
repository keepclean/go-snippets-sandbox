package main

import (
	"fmt"
	"sync"
	"time"
)

func aaaa() {
	var wg sync.WaitGroup
	var counter = 10
	var startTime = time.Now()

	for i := -1; i <= counter; i++ {
		wg.Add(1)
		go func(x int) {
			fmt.Println("i:", x)
			time.Sleep(time.Duration(x) * time.Second)
			fmt.Println("I slept:", x, "seconds")
			wg.Done()
		}(i)
	}
	wg.Wait()

	t := time.Now()
	fmt.Println("Back to main")
	fmt.Println("Total time:", t.Sub(startTime))
}
