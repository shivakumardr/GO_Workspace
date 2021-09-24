package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{2, 5, 43, 2, 1, 4, 3, 7, 89, 6}
	xs := []string{"ka", "fa", "rd", "eq", "xc", "cx", "rr", "sw", "is", "fncx", "aa"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)
	fmt.Println("......................")
	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)
}
