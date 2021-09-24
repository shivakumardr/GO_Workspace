package main

import "fmt"

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)
	x++
	fmt.Println(x)
}
func foo() {
	fmt.Println(" I am in foo:", x)
	x++
}
