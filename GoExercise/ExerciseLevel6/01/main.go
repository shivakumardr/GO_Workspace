package main

import "fmt"

/*
create a func with the identifier foo that returns an int
create a func with the identifier bar that returns an int and a string
call both funcs
print out their results

*/
func main() {
	i := foo()
	fmt.Println(i)
	fmt.Println(foo())
	x, st := bar()
	fmt.Println(x, st)
}
func foo() int {
	return 42
}
func bar() (int, string) {
	return 42, "hell function"
}
