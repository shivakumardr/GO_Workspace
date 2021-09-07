package main

import "fmt"

/*
Hands-on exercise #6
Create a program that shows the “if statement” in action.
solution: https://play.golang.org/p/DpZ_FLfn5s
*/
func main() {
	st := "Om Nama"
	if st == "Om" {
		fmt.Println("It,s Om")
	} else if st == "Nama" {
		fmt.Println("It,s Nama")
	} else if st == "Om Nama" {
		fmt.Println("It,s Om Nama")
	} else {
		fmt.Println("Nothing is matching")
	}
}
