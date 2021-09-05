package main

import "fmt"

/*
Hands-on exercise #2
Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.

*/
func main() {
	a := 42 == 42
	b := 42 <= 43
	c := 43 >= 42
	d := 43 != 42
	e := 43 > 42
	f := 43 < 44
	fmt.Println(a, b, c, d, e, f)
}
