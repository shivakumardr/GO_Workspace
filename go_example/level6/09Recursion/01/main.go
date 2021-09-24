package main

import "fmt"

func main() {
	fmt.Println(4 * 3 * 2 * 1)
	n := factrorial(9)
	fmt.Println(n)
	r := recursion(9)
	fmt.Println(r)
	l := loopfact(9)
	fmt.Println(l)
}
func factrorial(n int) int {
	var fact int
	fact = n
	n = n - 1
	for n > 0 {
		fact *= n
		n = n - 1
	}
	return fact
}
func recursion(n int) int {
	if n == 1 {
		return 1
	}
	return n * recursion(n-1)
}
func loopfact(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
