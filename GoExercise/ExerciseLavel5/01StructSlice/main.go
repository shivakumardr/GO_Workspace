package main

/*
Hands-on exercise #1
Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
first name
last name
favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.
solution:
https://play.golang.org/p/ouRHmH_POg
*/
import "fmt"

type person struct {
	firstName  string
	lastName   string
	favFlavors []string
}

func main() {
	p1 := person{
		firstName: "Shiva",
		lastName:  "Kumar",
		favFlavors: []string{
			"Vanilla",
			"Strawberry",
			"chacolate",
		},
	}
	p2 := person{
		firstName: "Pushpa",
		lastName:  "M",
		favFlavors: []string{
			"Buttered Pecan",
			"Strawberry",
			"chacolate",
		},
	}

	fmt.Println(p1.firstName, p1.lastName)
	//	fmt.Println(p1.lastName)
	for i, v := range p1.favFlavors {
		fmt.Println("\t", i, v)
	}

	fmt.Println(p2.firstName, p2.lastName)
	//	fmt.Println(p2.lastName)
	for i, v := range p2.favFlavors {
		fmt.Println("\t", i, v)
	}
}
