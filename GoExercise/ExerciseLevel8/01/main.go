package main

import (
	"encoding/json"
	"fmt"
)

/*
Hands-on exercise #1
Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.
solution: https://play.golang.org/p/8BK6PXj3aG
*/
type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Shiva",
		Age:   37,
	}
	u2 := user{"Pushpa", 34}
	u3 := user{"pavaki", 6}
	u4 := user{"Parvathi", 6}
	people := []user{u1, u2, u3, u4}
	fmt.Println(people)
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
