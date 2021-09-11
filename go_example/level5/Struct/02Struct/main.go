package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}
type secrateAgent struct {
	person
	firstName    string
	licenstoKill bool
}

func main() {
	sa1 := secrateAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       32,
		},
		firstName:    "JamesBond",
		licenstoKill: true,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.person.firstName, sa1.lastName, sa1.age, sa1.firstName, sa1.licenstoKill)
}
