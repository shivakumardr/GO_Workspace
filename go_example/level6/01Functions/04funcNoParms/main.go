package main

import "fmt"

func main() {
//	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
//	sum := sum(xi...)
	sum := sum()
	fmt.Println("Total sum of n number is:", sum)
}
func sum(x ...int) int {
	totalsum := 0
	for i, val := range x {
		totalsum += val
		fmt.Println("for index i", i, "Sum of the value ", val, "is totalsum ", totalsum)
	}
	return totalsum
}

