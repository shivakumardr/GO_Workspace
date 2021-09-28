package main

import "fmt"

func main() {
	x := []int{1, 2, 3}
	fmt.Println(x)
	fmt.Println("..............")
	x = append(x, 4, 5, 6)
	fmt.Println(x)
	fmt.Println("..............")
	y := []int{21, 34, 24, 44}
	x = append(x, y...)
	fmt.Println(x)
	fmt.Println("..............")

	x = append(x[:2], x[4:]...)
	fmt.Println(x)
}
