package main

import "fmt"

/*
Hands-on exercise #2
create a person struct
create a func called “changeMe” with a *person as a parameter
change the value stored at the *person address
important
to dereference a struct, use (*value).field
p1.first
(*p1).first
“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
https://golang.org/ref/spec#Selectors
create a value of type person
print out the value
in func main
call “changeMe”
in func main
print out the value
code: https://play.golang.org/p/AehM662HKS
*/
type person struct {
	name string
}

func main() {
	p1 := person{
		name: "James Bond 007",
	}
	fmt.Println(p1)
	fmt.Println(&p1.name)
	changeMe(&p1)
	fmt.Println(p1)
	fmt.Println(&p1)
	fmt.Printf("%T\n", p1)
}
func changeMe(p *person) {
	//p.name = "Shivakumar D R"
	(*p).name = "Shivakumar D R" // both works
}
