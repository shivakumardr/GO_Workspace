package main

import "fmt"

/*
create a func with the identifier foo that
takes in a variadic parameter of type int
pass in a value of type []int into your func (unfurl the []int)
returns the sum of all values of type int passed in
create a func with the identifier bar that
takes in a parameter of type []int
returns the sum of all values of type int passed in
code: https://play.golang.org/p/B0yRxtBQPD
*/
func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	si := foo(xi...)
	fmt.Println("Sum of foo func", si)
	zi := []int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	si = bar(zi)
	fmt.Println("Sum of bar func", si)
}
func foo(x ...int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
