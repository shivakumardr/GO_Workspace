package main

import "fmt"

func main() {
	foo()
	// Anonyumous functions
	func() {
		fmt.Println("I am in anonyumous func")
	}()
	func(x int) {
		fmt.Println("I am in anonyumous parmeter func", x)
	}(42)
}
func foo() {
	fmt.Println("I am in fooo")
}
