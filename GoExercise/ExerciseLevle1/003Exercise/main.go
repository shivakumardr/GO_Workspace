package main

import "fmt"

/*
Using the code from the previous exercise,
At the package level scope, assign the following values to the three variables
for x assign 42
for y assign “James Bond”
for z assign true
in func main
use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
print out the value stored by variable “s”

*/
var x int
var y string
var z bool

func main() {
	x = 42
	y = "All is well"
	z = true
	s := fmt.Sprintf("%v,\t%v,\t%v", x, y, z)
	fmt.Println(s)

}
