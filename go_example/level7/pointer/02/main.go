package main

import "fmt"

func main() {
	i := 0
	fmt.Println("Value in main func before call", i)
	foo(i)
	fmt.Println("Value in main func after calling foo", i)
	bar(&i)
	fmt.Println("Value in main func after calling bar", i)
}
func foo(x int) {
	fmt.Println(x)
	x = 40
	fmt.Println(x)
}
func bar(y *int) {
	fmt.Println(*y)
	*y = 70
	fmt.Println(*y)
}
