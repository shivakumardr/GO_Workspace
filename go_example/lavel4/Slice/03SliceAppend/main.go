package main

import "fmt"
func main(){
	x:= []int{1,2,3}
	fmt.Println(x)
	x = append(x, 4,5,6)
	fmt.Println(x)
	y:= []int {21,34,24,44}
	x = append(x, y...)
	fmt.Println(x)
}
