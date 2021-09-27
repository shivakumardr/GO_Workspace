package main

import (
	"fmt"
	"runtime"
	"sync"
)
func main(){
	fmt.Println("CPUs",runtime.NumCPU())
	fmt.Println("Go routines",runtime.NumGoroutine())
	count:=0
	const gs =100
	var wg sync.WaitGroup
	wg.Add(gs)
	for i:=0;i<gs;i++{
		go func(){
			v:=count
			runtime.Gosched()
			v++
			count=v
			wg.Done()
		}()
		fmt.Println("Go routines",runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Go routines",runtime.NumGoroutine())
	fmt.Println("Counter:",count)

}
