package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin cpus", runtime.NumCPU())
	fmt.Println("Begin GO routiens ", runtime.NumGoroutine())
	cnt := 0
	const gs = 100
	var wg sync.WaitGroup
	var mut sync.Mutex
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			mut.Lock()
			v := cnt
			v++
			cnt = v
			fmt.Println("Go Count", cnt)
			mut.Unlock()
			wg.Done()
		}()
		fmt.Println("Mid cpus", runtime.NumCPU())
		fmt.Println("Mid GO routiens ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Going to finish")
	fmt.Println("End cpus", runtime.NumCPU())
	fmt.Println("End GO routiens ", runtime.NumGoroutine())
	fmt.Println("Count", cnt)

}
