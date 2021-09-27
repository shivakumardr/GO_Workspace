package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs", runtime.NumCPU())
	fmt.Println("Go routines", runtime.NumGoroutine())
	var count int64
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			runtime.Gosched()
			fmt.Println("\t\tAtomic:", atomic.LoadInt64(&count))
			wg.Done()
		}()
		fmt.Println("Go routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines", runtime.NumGoroutine())
	fmt.Println("\t\tCounter:", count)

}
