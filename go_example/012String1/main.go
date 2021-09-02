package main

import "fmt"

func main() {
	s := "Wlecome to go Strings"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	// Print UTF8 code
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[i] %v, = %#U\n", s[i], s[i])
	}
	for i, v := range s {
		fmt.Printf("\nIndex of %v, is having hex value %#x", i, v)

	}

}
