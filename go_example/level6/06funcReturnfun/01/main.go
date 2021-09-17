package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)
	// this func will return value
	x := func() int {
		return 452
	}()
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	// This func will return the address of value
	y := func() int {
		return 453
	}
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	//  this is fun returning func
	z := bar()
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	i := z()
	fmt.Println(i)
	fmt.Printf("%T\n", i)

}
func foo() int {
	return 10
}
func bar() func() int {
	return func() int {
		return 451
	}
}

