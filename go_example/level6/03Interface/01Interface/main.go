package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}
type secrateAgent struct {
	person
	ltk bool
}

func (s secrateAgent) speak() {
	fmt.Println("i am", s.first, s.last, s.ltk, "-  the secrateAgent speak")
}
func (p person) speak() {
	fmt.Println("i am", p.first, p.last, "- The person speak")
}

type human interface {
	speak()
}

func bar(h human) {
	switch h.(type) {// assering
	case person:
		fmt.Println("I was passed in to barrrrr", h.(person).first)
	case secrateAgent:
		fmt.Println("I was passed in to barrrrr", h.(secrateAgent).first)
	default:
		fmt.Println("I was passed in to bar", h)

	}
	fmt.Println("I was passed in to bar", h)
}

// conversion
type hotdog int

func main() {
	sa1 := secrateAgent{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	sa2 := secrateAgent{
		person: person{
			"Money",
			"Penny",
		},
		ltk: true,
	}
	p1 := person{
		"Shiva",
		"Kumar",
	}
	fmt.Println(sa1)
	fmt.Println(sa2)
	fmt.Println(p1)
	sa1.speak()
	sa2.speak()
	bar(sa1)
	bar(sa2)
	bar(p1)
	// converstion
	var x hotdog
	x = 42
	fmt.Println(x)
	fmt.Printf("%T", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)

}
