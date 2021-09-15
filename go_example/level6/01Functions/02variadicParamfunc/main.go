package main

import "fmt"

func main() {
	sum := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
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
