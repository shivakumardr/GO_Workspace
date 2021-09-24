package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	p1 := person{
		FirstName: "James",
		LastName:  "Bond",
		Age:       37,
	}
	p2 := person{
		FirstName: "Miss",
		LastName:  "MoneyPenny",
		Age:       27,
	}
	pepole := []person{p1, p2}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(pepole)
	bs, err := json.Marshal(pepole)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(bs)
	fmt.Println(string(bs))
}
