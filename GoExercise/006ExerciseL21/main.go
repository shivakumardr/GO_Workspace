package main

import "fmt"

/*
Hands-on exercise #1
Write a program that prints a number in decimal, binary, and hex
*/
func main() {
	x := 20
	fmt.Println(x)
	fmt.Printf("Decimal=%d\tBinary=%b\tHexadecimal=%#x", x, x, x)
}
