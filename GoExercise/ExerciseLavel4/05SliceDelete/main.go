package main

/*
Hands-on exercise #5
To DELETE from a slice, we use APPEND along with SLICING. For this hands-on exercise, follow these steps:
start with this slice
x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
use APPEND & SLICING to get these values here which you should ASSIGN to a variable “y” and then print:
[42, 43, 44, 48, 49, 50, 51]
solution: https://play.golang.org/p/u8zpHLfgSE
just for fun:
https://goo.gl/frz786
https://www.youtube.com/channel/UCJ8YIwWQCO7hMiqpOw2ZLFw
*/

import "fmt"

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x)
	y := append(x[:3], x[6:]...)
	fmt.Println(y)

}
