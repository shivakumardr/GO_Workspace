package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("I am in func experssion")
	}
	f()
	g := func(x int) {
		fmt.Println("I am in Parameterised fun experssion:", x)
	}
	g(32)

}
