package main

import "fmt"

/*
Hands-on exercise #7
Assign a func to a variable, then call that func
code: https://play.golang.org/p/_Qu7ZAyFDH
*/
var x int
var g func() = func() {
	fmt.Println(" g calling outside func")
}

func main() {
	f := func() {
		for i := 1; i < 10; i++ {
			if (i % 4) == 0 {
				fmt.Println("Number", i, " is div by 4 between 1 to 100")
			}
		}
	}
	f()
	fmt.Printf("This of type:%T\n", f)
	fmt.Println(x)
	fmt.Printf("This of type:%T\n", x)
	g()
	g = f
	g()
	fmt.Printf("This is g of type:%T\n", g)
}
