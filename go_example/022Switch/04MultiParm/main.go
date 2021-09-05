package main

import "fmt"

func main() {
	n := "OmNama"
	switch n {
	case "shiva":
		fmt.Println("this is shiva")
	case "g", "OmNama", "M":
		fmt.Println("This prints either of g or OmNama or M")
		fallthrough
	default:
		fmt.Println("This is default")

	}
}
