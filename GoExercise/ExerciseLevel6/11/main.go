package main

import "fmt"

func main() {
	fact := factorial(4)
	fmt.Println(fact)
	fact = factorial1(4)
	fmt.Println(fact)
	fact = factorial2(4)
	fmt.Println(fact)
}
func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)

}
func factorial1(n int) int {
	f := 1
	for ; n > 0; n-- {
		f *= n
	}
	return f
}
func factorial2(n int) int {
	var fact int
	fact = n
	n = n - 1
	for n > 0 {
		fact *= n
		n = n - 1
	}
	return fact
}
