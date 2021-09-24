package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sm := sum(ii...)
	fmt.Println("Sum of all number:", sm)

	sm = even(sum, ii...)
	fmt.Println("Sum of even number:", sm)
	sm = odd(sum, ii...)
	fmt.Println("Sum of odd number:", sm)

}
func sum(xi ...int) int {
	//	fmt.Printf("%T\n", xi)
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 1 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
