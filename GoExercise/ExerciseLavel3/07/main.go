package main

import "fmt"

/*
Hands-on exercise #7
Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
solution: https://play.golang.org/p/IDnrJpE7vT
*/
func main() {
	st := ""
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
