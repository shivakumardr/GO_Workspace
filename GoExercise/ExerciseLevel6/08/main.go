package main

import "fmt"

/*
Hands-on exercise #8
Create a func which returns a func
assign the returned func to a variable
call the returned func
code: https://play.golang.org/p/c2HwqVE1Rd
*/
func main() {
	f := foo()
	fmt.Println(f())
}
func foo() func() int {
	return func() int {
		return 42
	}
}
