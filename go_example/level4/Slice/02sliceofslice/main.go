package main

import "fmt"

// Slice of Slice
func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(x)
	fmt.Println(".............")
	fmt.Println(len(x))
	fmt.Println(".............")
	fmt.Println(x[1:])
	fmt.Println(x[3:6])
	fmt.Println(x[1:4])
	fmt.Println(".............")
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Println(".............")
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}
