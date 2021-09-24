package main

import "fmt"

/*
Hands-on exercise #6
Build and use an anonymous func
code: https://play.golang.org/p/DQX3xEIcRe

*/
func main() {
	foo()
	func() {
		for i := 0; i <= 100; i++ {
			fmt.Println(i)
		}
		fmt.Println("done with anonyumous func")
	}()
	func(x int) {
		fmt.Println(x)
	}(43)
	fmt.Println("done with anonyumous param func")
}
func foo() {
	fmt.Println("i am in foo")
}
