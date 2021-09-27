package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Begin CPUs", runtime.NumCPU())
	fmt.Println("Begin Goroutiens", runtime.NumGoroutine())

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		fmt.Println("I am in go first fun ")
		wg.Done()
	}()
	go func() {
		fmt.Println("i am in go second func")
		wg.Done()
	}()
	fmt.Println("mid CPUs", runtime.NumCPU())
	fmt.Println("mid Goroutiens", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("About to exit")
	fmt.Println("Final CPUs", runtime.NumCPU())
	fmt.Println("Final Goroutiens", runtime.NumGoroutine())

}
