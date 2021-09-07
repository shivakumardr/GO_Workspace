package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	a := 7
	b := 32
	c := 7
	fmt.Println(a == b)
	fmt.Println(a == c)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(c >= b)
	fmt.Println(c >= a)

}
