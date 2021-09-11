package main

import "fmt"

/*
Hands-on exercise #1
Using a COMPOSITE LITERAL:
create an ARRAY which holds 5 VALUES of TYPE int
assign VALUES to each index position.
Range over the array and print the values out.
Using format printing
print out the TYPE of the array

*/

func main() {
	Arr := [5]int{1, 2, 3, 4, 5}
	fmt.Println(Arr)
	for i := 0; i < len(Arr); i++ {
		fmt.Println(i, Arr[i])
	}
	fmt.Println("..............")
	for i, v := range Arr {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", Arr)
}
