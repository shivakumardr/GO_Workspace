package main

import (
	"fmt"
)

/*
Hands-on exercise #8
Create a program that uses a switch statement with no switch expression specified.
solution: https://play.golang.org/p/YpPgkWGqKY

*/

func main() {
	switch {
	case false:
		fmt.Println("should not print")
	case true:
		fmt.Println("should print")
	}
}
