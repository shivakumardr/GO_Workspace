package main

import "fmt"

/*
Hands-on exercise #2
Using a COMPOSITE LITERAL:
create a SLICE of TYPE int
assign 10 VALUES
Range over the slice and print the values out.
Using format printing
print out the TYPE of the slice

*/
func main() {
	xSlice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(xSlice)
	for i, v := range xSlice {
		fmt.Println(i, v)
		// fmt.Printf("%T\n", v)
	}
	fmt.Printf("%T\n", xSlice)
}
