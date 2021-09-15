package main

import "fmt"

func main() {
	fmt.Println("I am in Main")
	defer foo()
	bar()
}
func foo() {
	fmt.Println("I am in Foo")
	fmt.Println("Defer  func call when end of prg } encounter")
}
func bar() {
	fmt.Println("I am in bar")
}
