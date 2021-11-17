package main

import "fmt"

func main() {
	fmt.Println("sum of 2 + 3 = ", mySum(2, 3))
	fmt.Println("sum of 4 + 5 = ", mySum(4, 5))
	fmt.Println("sum of 6 + 5 = ", mySum(6, 5))
}
func mySum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum + 1
}
