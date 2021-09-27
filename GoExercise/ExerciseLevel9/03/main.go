package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin Cpus", runtime.NumCPU())
	fmt.Println("Begin Go Routiens", runtime.NumGoroutine())
	cnt := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := cnt
			runtime.Gosched()
			v++
			cnt = v
			wg.Done()
		}()
		fmt.Println("mid Go routines", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Mid go routiens", runtime.NumGoroutine())
	fmt.Println("Count", cnt)
}
