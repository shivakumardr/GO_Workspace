                package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	b := &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", &b)

	fmt.Println(*&a)
	fmt.Printf("%T\n", *&a)
	bb := &b
	fmt.Println(bb)
	fmt.Println(*bb)
	fmt.Println(**bb)
	fmt.Printf("%T\n", bb)
	fmt.Printf("%T\n", *bb)
	**bb = 44
	fmt.Println(a)
}
