package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	p1 := person{
		firstName: "Shiva",
		lastName:  "Kumar",
		age:       38,
	}
	p2 := person{
		firstName: "Pushpa",
		lastName:  "M",
		age:       34,
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("Please enter the details of persons")
//	fmt.Scanf("%v,%v,%v", p1.firstName, p1.lastName, p1.age)
	fmt.Println(p1.firstName, p1.lastName, p1.age)
	fmt.Println(p2.firstName, p2.lastName, p2.age)
}
