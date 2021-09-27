package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t", runtime.GOARCH)
	fmt.Println("Num CPU\t\t", runtime.NumCPU())
	fmt.Println("Go routines\t", runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Num CPU\t\t", runtime.NumCPU())
	fmt.Println("Go routines\t", runtime.NumGoroutine())
	wg.Wait()
}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("Bar", i)
	}
}
