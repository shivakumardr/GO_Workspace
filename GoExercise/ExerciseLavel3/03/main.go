package main

import "fmt"

/*
Hands-on exercise #3
Create a for loop using this syntax
for condition { }
Have it print out the years you have been alive.
solution: https://play.golang.org/p/tnyqBPJ-i5

*/
func main() {
	bd := 1983

	for bd < 2022 {
		fmt.Println(bd, bd-1982)
		bd++
	}
}
