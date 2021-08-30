package main

import "fmt"

var b int
type myInt int
var a myInt
func main(){
	b= 42
	a = 82
	fmt.Println(b)
	fmt.Printf("%T\n", b)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	b = int(a)
	fmt.Println(a)
	fmt.Printf("%T\n", a)


}
