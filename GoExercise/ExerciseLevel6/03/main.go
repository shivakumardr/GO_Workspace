package main

import "fmt"

/*
Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
code: https://play.golang.org/p/XluEuUD0Nw
*/
func main() {
	fmt.Println("i am in func main")
	defer foo()
	bar()
}
func foo() {
	fmt.Println("i am in foo")
}
func bar() {
	defer func() {
		fmt.Println("i am in defer bar")
	}()
	fmt.Println("i am in bar")
}
