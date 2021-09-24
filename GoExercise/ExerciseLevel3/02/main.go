package main

import "fmt"

/*
Hands-on exercise #2
Print every rune code point of the uppercase alphabet three times. Your output should look like this:
65
	U+0041 'A'
	U+0041 'A'
U+0041 'A'
66
	U+0042 'B'
	U+0042 'B'
	U+0042 'B'
 … through the rest of the alphabet characters
solution: https://play.golang.org/p/1OjnCX1D5H

*/
func main() {
	x := 65
	for i := 0; i < 26; i++ {
		fmt.Printf("......'%c'..........\n", x+i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%v\t%#U\n", x+i, x+i)
		}

	}

}
