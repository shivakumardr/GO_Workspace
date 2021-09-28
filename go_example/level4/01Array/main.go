package main

import (
	"fmt"
)

// Array is mainly not used in GO
func main() {
	var x [5]int
	var y [7]float64
	fmt.Println(x)
	fmt.Println(y)
	y[5] = 4.8787
	x[3] = 900
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(len(x))
	fmt.Println(len(y))

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
