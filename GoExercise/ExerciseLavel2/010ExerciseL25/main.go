package main

import "fmt"
/*
Hands-on exercise #5
Create a variable of type string using a raw string literal. Print it.
 */
func main() {
	a := `Here is somthing,
as a
Raw string,
"We are going to print it"
`
	fmt.Println(a)
}
