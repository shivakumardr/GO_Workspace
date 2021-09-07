package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("............")
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5])
	fmt.Println("............")
	for i, v:= range x{
		fmt.Println(i,v)
	}
}
