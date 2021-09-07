package main

import "fmt"

/*
understanding zero value
false for booleans
0 for integers
0.0 for floats
"" for strings
nil for
pointers
functions
interfaces
slices
channels
maps
use short declaration operator as much as possible
use var for
zero value
package scope

*/
var z string
var x int
var y float32
var xx complex64

func main() {
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(xx)
	fmt.Printf("%T\n", xx)
}
