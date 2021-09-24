package main

import "fmt"
/*
Create a value and assign it to a variable.
Print the address of that value.
code: https://play.golang.org/p/57kW8xd0qT

 */
func main() {
	x := 42
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Printf("Type of x:%T\n", x)
}
