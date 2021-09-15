package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
	age   int
}
type SecrateAgent struct {
	person
	ltk bool
}

func (s SecrateAgent) speak() {
	fmt.Println("First Name", s.first)
	fmt.Println("last Name", s.last)
	fmt.Println("Age", s.age)
	fmt.Println("License to kill", s.ltk)
	fmt.Println(">>>>.........<<<<<<<")
}
func main() {
	sa1 := SecrateAgent{
		person: person{
			first: "james",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	sa2 := SecrateAgent{
		person: person{
			"Miss",
			"MoneyPenny",
			27,
		},
		ltk: false,
	}
	sa1.speak()
	sa2.speak()
}
